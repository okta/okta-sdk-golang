/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

const _ = require('lodash');
_.mixin(require('lodash-inflection'));
const fs = require('fs');

const golang = module.exports;

function getType(obj, prefix = "") {
  switch (obj.commonType) {
    case 'dateTime' :
      return String.raw`*time.Time`;
    case 'integer' :
      return String.raw`int64`;
    case 'boolean' :
      return String.raw`*bool`;
    case 'hash' :
      if (obj.model === 'object') {
        return String.raw`interface{}`;
      }
      if (obj.model === 'boolean') {
        return String.raw`map[string]bool`;
      }
      return String.raw`map[string]*` + obj.model;
    case 'array' :
      if (obj.model === undefined || obj.model === "string") {
        return String.raw`[]string`;
      } else {
        return String.raw`[]` + prefix + obj.model;
      }
    case 'enum' :
    case '':
    case 'null' :
    case 'password' :
    case 'binary' :
      return String.raw`string`;
    case 'double':
      return String.raw`float64`;
    case 'object' :
      // edge case from org settings
      if (obj.propertyName === "_links") {
        return "interface{}";
      }
      if (obj.model === "UserSchemaBaseProperties") {
        return String.raw`map[string]*UserSchemaAttribute`;
      }
      if (obj.model === "GroupSchemaBaseProperties") {
        return String.raw`map[string]*GroupSchemaAttribute`;
      }
      if (obj.model === undefined) {
        return String.raw`string`;
      } else {
        if (prefix !== "") {
          return prefix + obj.model;
        }
        return obj.model;
      }
    default:
      if (obj.propertyName === "pattern" || obj.propertyName === "admin" || obj.propertyName === "enduser") {
        return String.raw`*string`;
      }
      return obj.commonType;
  }
}

function lowercaseFirstLetter(text) {
  // Only works for standard english characters
  const isSingleCap = /^[A-Z][^A-Z]/;
  const isAllCap = /^[A-Z]*$/;
  const isMultiCap = /^[A-Z]{2,}/;
  if (text.match(isSingleCap)) {
    return text.charAt(0).toLowerCase() + text.slice(1);
  } else if (text.match(isAllCap)) {
    return text.toLowerCase();
  } else if (text.match(isMultiCap)) {
    return text.replace(/^([A-Z]*)([A-Z])/, function (match, leading, keep, offset, remaining) {
      return leading.toLowerCase() + keep;
    });
  } else {
    return text;
  }
}

function ucFirst(text) {
  switch (text) {
    case 'csr':
      return 'Csr';
    case 'csrMetadata':
      return 'CsrMetadata';
    case 'csrMetadataSubject':
      return 'CsrMetadataSubject';
    case 'csrMetadataSubjectAltNames':
      return 'CsrMetadataSubjectAllNames';
    default:
      return text.charAt(0).toUpperCase() + text.slice(1);
  }

}

function strToUpper(string) {
  return string.toUpperCase(string)
}

function structProp(prop) {
  prop = prop.replace(/#/g, "");
  prop = prop.replace(/\$/g, "");
  prop = prop.replace(/(_\w)/g, function (m) {
    return m[1].toUpperCase();
  });

  prop = prop.charAt(0).toUpperCase() + prop.slice(1);

  return prop;
}

function getImports(object) {
  let imports = [];

  if (object.model.properties !== undefined) {
    for (let property of object.model.properties) {
      switch (property.commonType) {
        case 'dateTime' :
          imports.push("time");
      }
    }
  }
  if (object.model.parent !== undefined) {
    for (let property of object.model.parent.properties) {
      switch (property.commonType) {
        case 'dateTime' :
          imports.push("time");
      }
    }

    if (object.model.parent.parent !== undefined) {
      for (let property of object.model.parent.parent.properties) {
        switch (property.commonType) {
          case 'dateTime' :
            imports.push("time");
        }
      }
    }
  }

  if (object.operations !== undefined) {
    for (let operation in object.operations) {
      if (object.operations[operation].queryParams.length) {
        imports.push("github.com/okta/okta-sdk-golang/v2/okta/query")
        imports.push("context");
      }
    }
  }

  if (object.model.methods !== undefined) {
    for (let method of object.model.methods) {

      if (method.operation.queryParams.length) {
        imports.push("github.com/okta/okta-sdk-golang/v2/okta/query")
      }
      imports.push("fmt");
      imports.push("context");
      if (method.operation.responseModel !== undefined) {
        imports.push("fmt");
      }

      if (method.operation.bodyModel !== undefined) {
        imports.push("fmt");
      }
    }
  }

  if (object.model.crud !== undefined) {
    for (let method of object.model.crud) {
      if (method.operation.queryParams.length) {
        imports.push("github.com/okta/okta-sdk-golang/v2/okta/query");
      }
      imports.push("fmt");
      imports.push("context");
      if (method.operation.responseModel !== undefined) {
        imports.push("fmt");
      }

      if (method.operation.bodyModel !== undefined) {
        imports.push("fmt");
      }
    }
  }

  if (object.model.modelName === "LogEvent") {
    imports.push("fmt");
  }

  if (object.model.modelName === "UserSchema") {
    imports.push("fmt");
    imports.push("context");
  }

  if (object.model.modelName === "Domain") {
    imports.push("fmt");
    imports.push("context");
  }

  if (object.model.modelName === "DomainCertificate") {
    imports = [];
  }

  imports = [...new Set(imports)];

  if (object.model.modelName === "Role") {
    imports.splice(imports.indexOf("context"), 1);
    imports.splice(imports.indexOf("fmt"), 1);
  }

  return imports;
}

function operationArgumentBuilder(operation) {
  const args = [];

  args.push("ctx context.Context");

  operation.pathParams.map((arg) => args.push(arg.name + " " + arg.type));

  if ((operation.method === 'post' || operation.method === 'put') && operation.bodyModel) {
    let bodyModel = ucFirst(_.camelCase(operation.bodyModel));

    if (bodyModel === "Application") {
      bodyModel = "App";
    }

    if (bodyModel === "UserFactor") {
      bodyModel = "Factor";
    }

    if (bodyModel === "String") {
      bodyModel = "string";
    }

    args.push(`body ` + bodyModel);
  }

  if (operation.operationId === "getApplication") {
    args.push(`appInstance App`);
  }

  if (operation.operationId === "getFactor" ||
    operation.operationId === "activateFactor" ||
    operation.operationId === "verifyFactor") {
    args.push(`factorInstance Factor`);
  }

  if (operation.queryParams.length) {
    args.push('qp *query.Params');
  }

  return args.join(', ');
}

function getPath(operation) {
  let path = operation.path;
  for (let param of operation.pathParams) {
    path = path.replace(`{${param.name}}`, String.raw`%v`);
  }

  return `"${path}"`;
}

function getPathParams(operation) {
  const args = []
  for (let param of operation.pathParams) {
    args.push(param.name)
  }
  return args.join(', ');
}

function returnType(operation) {
  if (operation.responseModel !== undefined) {
    let responseModel = "*" + operation.responseModel
    if (responseModel === "*Application") {
      if (applicationModelInterface) {
        responseModel = "App"
      } else {
        responseModel = "interface{}"
      }
    }
    if (responseModel === "*UserFactor") {
      if (factorModelInterface) {
        responseModel = "Factor"
      } else {
        responseModel = "interface{}"
      }
    }
    if (operation.isArray !== undefined && operation.isArray === true) {
      return " ([]" + responseModel + ", *Response, error) ";
    }
    return " (" + responseModel + ", *Response, error) ";
  }

  return " (*Response, error) ";
}

function getClientTags(operations) {
  let tags = []
  for (let operation of operations) {

    tags.push(operation.tags[0]);
  }

  tags = [...new Set(tags)];

  return tags;

}

function responseModelInterface(operationId) {
  return operationId === "listFactors" ||
    operationId === "listSupportedFactors" ||
    operationId === "listApplications" ||
    operationId === "listAppTargetsForRole" ||
    operationId === "listApplicationTargetsForApplicationAdministratorRoleForGroup" ||
    operationId === "listApplicationTargetsForApplicationAdministratorRoleForUser" ||
    operationId === "listAssignedApplicationsForGroup" ||
    operationId === "listAssignedApplicationsForUser";
}

function applicationModelInterface(operationId) {
  return operationId === "listApplications" ||
    operationId === "listAppTargetsForRole" ||
    operationId === "listAssignedApplicationsForGroup" ||
    operationId === "listAssignedApplicationsForUser";
}

function factorModelInterface(operationId) {
  return operationId === "listFactors" ||
    operationId === "listSupportedFactors";
}

function catalogApplicationInterface(operationId) {
  return operationId === "listApplicationTargetsForApplicationAdministratorRoleForGroup" ||
    operationId === "listApplicationTargetsForApplicationAdministratorRoleForUser";
}

function factorInstanceOperation(operationId) {
  return operationId === "getFactor" ||
    operationId === "activateFactor" ||
    operationId === "verifyFactor";
}

function getClientTagResources(operations) {
  let tags = getClientTags(operations);
  let tagResources = []
  for (let tag of tags) {
    if (tag === "AuthServer") tag = "AuthorizationServer";
    if (tag === "Template") tag = "SmsTemplate";
    if (tag === "Idp") tag = "IdpTrust";
    if (tag === "UserFactor") tag = "UserFactor";
    if (tag === "Log") tag = "LogEvent";
    if (tag === "Org") tag = "OrgSetting";
    if (tag === "ThreatInsight") tag = "ThreatInsightConfiguration";
    tagResources.push(structProp(tag) + " *" + structProp(tag) + "Resource")
  }
  tagResources.sort();
  return tagResources.join("\n\t");
}

function getNewClientTagProps(operations) {
  let tags = getClientTags(operations);
  let tagResources = []
  for (let tag of tags) {
    if (tag === "AuthServer") tag = "AuthorizationServer";
    if (tag === "Template") tag = "SmsTemplate";
    if (tag === "Idp") tag = "IdpTrust";
    if (tag === "UserFactor") tag = "UserFactor";
    if (tag === "Log") tag = "LogEvent";
    if (tag === "Org") tag = "OrgSetting";
    if (tag === "ThreatInsight") tag = "ThreatInsightConfiguration";
    tagResources.push("c." + structProp(tag) + " = (*" + structProp(tag) + "Resource)(&c.resource)")
  }
  tagResources.sort();
  return tagResources.join("\n\t");
}

function buildModelProperties(model) {
  const properties = {};
  const finalProps = [];

  if (model.parent !== undefined) {
    for (let parentProperty of model.parent.properties) {
      properties[parentProperty.propertyName] = parentProperty;
    }
    if (model.parent.parent !== undefined) {
      for (let parentProperty of model.parent.parent.properties) {
        properties[parentProperty.propertyName] = parentProperty;
      }
    }
  }

  if (model.properties !== undefined) {
    for (let modelProperty of model.properties) {
      properties[modelProperty.propertyName] = modelProperty;
    }
  }

  for (let propKey in properties) {
    finalProps.push(structProp(properties[propKey].propertyName) + " " +
      getType(properties[propKey], "*") + createJsonTag(properties[propKey].propertyName));
  }

  return finalProps.join("\n\t");
}

function createJsonTag(propertyName) {
  if (propertyName === "tokenLifetimeMinutes" ||
    propertyName === "accessTokenLifetimeMinutes" ||
    propertyName === "minLowerCase" ||
    propertyName === "minUpperCase" ||
    propertyName === "minNumber" ||
    propertyName === "minSymbol" ||
    propertyName === "maxSessionLifetimeMinutes" ||
    propertyName === "refreshTokenLifetimeMinutes" ||
    propertyName === "refreshTokenWindowMinutes" ||
    propertyName === "default_scope" ||
    propertyName === "userName" ||
    propertyName === "leeway" ||
    propertyName === "audienceOverride" ||
    propertyName === "defaultRelayState" ||
    propertyName === "destinationOverride" ||
    propertyName === "recipientOverride" ||
    propertyName === "ssoAcsUrlOverride" ||
    propertyName === "attributeStatements" ||
    propertyName === "admin" ||
    propertyName === "enduser" ||
    propertyName === "maxSessionIdleMinutes") {
    return " `json:\"" + propertyName + "\"`"
  } else {
    return " `json:\"" + propertyName + ",omitempty\"`"
  }
}

function isInstance(model) {
  if (model.modelName === "Csr" ||
    model.modelName === "CsrMetadata" ||
    model.modelName === "CsrMetadataSubject" ||
    model.modelName === "CsrMetadataSubjectAltNames" ||
    model.modelName === "OAuth2Claim" ||
    model.modelName === "OAuth2ScopeConsentGrant" ||
    model.modelName === "AcsEndpoint" ||
    model.modelName === "JwkUse" ||
    model.modelName === "OAuth2Actor" ||
    model.modelName === "OAuth2Client" ||
    model.modelName === "OAuth2RefreshToken" ||
    model.modelName === "OAuth2ClaimConditions" ||
    model.modelName === "OAuth2Scope" ||
    model.modelName === "WebAuthnUserFactorProfile" ||
    model.modelName === "OpenIdConnectApplicationSettingsClientKeys" ||
    model.modelName === "OpenIdConnectApplicationSettingsRefreshToken" ||
    model.modelName === "SingleLogout" ||
    model.modelName === "SpCertificate" ||
    model.modelName === "OpenIdConnectApplicationIdpInitiatedLogin" ||
    model.modelName === "ApplicationAccessibility" ||
    model.modelName === "ApplicationCredentials" ||
    model.modelName === "ApplicationCredentialsOAuthClient" ||
    model.modelName === "ApplicationCredentialsSigning" ||
    model.modelName === "ApplicationCredentialsUsernameTemplate" ||
    model.modelName === "ApplicationGroupAssignment" ||
    model.modelName === "ApplicationLicensing" ||
    model.modelName === "ApplicationSettings" ||
    model.modelName === "ApplicationSettingsApplication" ||
    model.modelName === "ApplicationSettingsNotes" ||
    model.modelName === "ApplicationSettingsNotifications" ||
    model.modelName === "ApplicationSettingsNotificationsVpn" ||
    model.modelName === "ApplicationSettingsNotificationsVpnNetwork" ||
    model.modelName === "ApplicationVisibility" ||
    model.modelName === "ApplicationVisibilityHide" ||
    model.modelName === "AppUser" ||
    model.modelName === "AppUserCredentials" ||
    model.modelName === "AppUserPasswordCredential" ||
    model.modelName === "AuthorizationServerCredentials" ||
    model.modelName === "AutoLoginApplicationSettings" ||
    model.modelName === "AutoLoginApplicationSettingsSignOn" ||
    model.modelName === "BasicApplicationSettings" ||
    model.modelName === "BasicApplicationSettingsApplication" ||
    model.modelName === "BookmarkApplicationSettings" ||
    model.modelName === "BookmarkApplicationSettingsApplication" ||
    model.modelName === "JsonWebKey" ||
    model.modelName === "OAuth2Token" ||
    model.modelName === "OAuthApplicationCredentials" ||
    model.modelName === "OpenIdConnectApplicationSettings" ||
    model.modelName === "OpenIdConnectApplicationSettingsClient" ||
    model.modelName === "SamlApplicationSettings" ||
    model.modelName === "SamlApplicationSettingsSignOn" ||
    model.modelName === "SamlAttributeStatement" ||
    model.modelName === "SchemeApplicationCredentials" ||
    model.modelName === "SecurePasswordStoreApplicationSettings" ||
    model.modelName === "SecurePasswordStoreApplicationSettingsApplication" ||
    model.modelName === "SignOnInlineHook" ||
    model.modelName === "SwaApplicationSettings" ||
    model.modelName === "SwaApplicationSettingsApplication" ||
    model.modelName === "SwaThreeFieldApplicationSettings" ||
    model.modelName === "SwaThreeFieldApplicationSettingsApplication" ||
    model.modelName === "WsFederationApplicationSettings" ||
    model.modelName === "WsFederationApplicationSettingsApplication" ||
    model.modelName === "OAuth2ScopesMediationPolicyRuleCondition") {
    return false
  }

  return model.tags[0] === "Application" || model.tags[0] === "UserFactor";
}

function log(item) {
  console.log(item);
}

golang.process = ({spec, operations, models, handlebars}) => {
  golang.spec = spec;
  const templates = [];
  const queryOptionsTemp = [];
  const queryOptions = [];
  const modelsByName = [];

  for (let model of models) {
    modelsByName[model.modelName] = model
  }

  for (let operation of operations) {
    for (let param of operation.queryParams) {
      queryOptionsTemp[param.name] = param.type;
    }
  }

  for (let key in queryOptionsTemp) {
    if (queryOptionsTemp[key] === "boolean") {
      queryOptionsTemp[key] = "bool";
    }
    if (queryOptionsTemp[key] === "integer") {
      queryOptionsTemp[key] = "int64";
    }
    queryOptions.push({name: key, type: queryOptionsTemp[key]});
  }

  templates.push({
    src: 'templates/query.go.hbs',
    dest: 'okta/query/query.go',
    context: {
      "queryOptions": queryOptions
    }
  });

  const version = process.env.OKTA_SDK_GOLANG_VERISON || spec.info.version;
  templates.push({
    src: 'templates/okta.go.hbs',
    dest: 'okta/okta.go',
    context: {
      "operations": operations,
      "models": models,
      "version": version
    }
  });

  for (let model of models) {

    if (model.extends !== undefined) {
      model.parent = modelsByName[model.extends];

      if (model.parent.resolutionStrategy !== undefined) {
        for (let value in model.parent.resolutionStrategy.valueToModelMapping) {
          if (model.modelName)
            if (model.parent.resolutionStrategy.valueToModelMapping[value] === model.modelName) {
              model.resolution = {fieldName: model.parent.resolutionStrategy.propertyName, fieldValue: value};
              let result = new Map()
              for (let value of model.properties) {
                result.set(value.propertyName, value)
              }
              for (let value of model.parent.properties) {
                if (!result.has(value.propertyName)) {
                  result.set(value.propertyName, value)
                }
              }
              model.properties = [...result].map(([name, value]) => (value));
            }
        }
      }

      if (model.parent.parent !== undefined && model.parent.parent.resolutionStrategy !== undefined) {
        model.resolution = model.parent.resolution;
      }
    }

    let modelOperations = {}

    if (model.crud !== undefined) {
      for (let modelCrud of model.crud) {
        modelOperations[modelCrud.operation.operationId] = modelCrud.operation;
      }
    }

    for (let operation of operations) {
      let tag = operation.tags[0];
      if (tag === "AuthServer") tag = "AuthorizationServer";
      if (tag === "Template") tag = "SmsTemplate";
      if (tag === "Idp") tag = "IdpTrust";
      if (tag === "UserFactor") tag = "UserFactor";
      if (tag === "Log") tag = "LogEvent";
      if (tag === model.modelName) {
        modelOperations[operation.operationId] = operation;
      }
    }

    templates.push({
      src: 'templates/model.go.hbs',
      dest: 'okta/' + lowercaseFirstLetter(model.modelName) + '.go',
      context: {
        "operations": modelOperations,
        "model": model
      }
    });
  }

  handlebars.registerHelper({
    getType,
    structProp,
    log,
    ucFirst,
    operationArgumentBuilder,
    getPath,
    getPathParams,
    returnType,
    getImports,
    strToUpper,
    lowercaseFirstLetter,
    getClientTagResources,
    getNewClientTagProps,
    buildModelProperties,
    responseModelInterface,
    applicationModelInterface,
    factorModelInterface,
    factorInstanceOperation,
    isInstance,
    catalogApplicationInterface
  });

  handlebars.registerPartial('partials.copyHeader', fs.readFileSync('generator/templates/partials/copyHeader.hbs', 'utf8'));
  handlebars.registerPartial('struct.withProp', fs.readFileSync('generator/templates/struct/withProp.go.hbs', 'utf8'));
  handlebars.registerPartial('model.imports', fs.readFileSync('generator/templates/model/imports.go.hbs', 'utf8'));
  handlebars.registerPartial('model.defaultMethod', fs.readFileSync('generator/templates/model/defaultMethod.go.hbs', 'utf8'));

  fs.writeFile("generator/createdFiles.json", JSON.stringify(templates), function (error) {
    console.log(error);
  });
  return templates;
};
