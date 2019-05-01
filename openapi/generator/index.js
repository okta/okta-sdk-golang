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

function getType(obj, prefix="") {
  switch(obj.commonType) {
    case 'dateTime' :
      return String.raw`*time.Time`;
    case 'integer' :
      return String.raw`int64`;
    case 'boolean' :
      return String.raw`*bool`;
    case 'hash' :
      return String.raw`interface{}`;
    case 'array' :
      return String.raw`[]string`;
    case 'enum' :
    case '':
    case 'null' :
    case 'password' :
      return String.raw`string`;
    case 'double':
      return String.raw`float64`;
    case 'object' :
      if(obj.model == undefined) {
        return String.raw`string`;
      } else {
        if(prefix !== null || prefix != "") {
          return prefix + obj.model;
        }
        return obj.model;
      }
    default:
      return obj.commonType;
  }
}

function lowercaseFirstLetter(string) {
  return string.charAt(0).toLowerCase() + string.slice(1);
}

function ucFirst(string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}

function strToUpper(string) {
  return string.toUpperCase(string)
}

function structProp(prop) {
  prop = prop.replace(/#/g,"");
  prop = prop.replace(/(\_\w)/g, function(m){return m[1].toUpperCase();});

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


  if (object.model.methods !== undefined) {
    for (let method of object.model.methods) {
      if(method.operation.queryParams.length) {
        imports.push("github.com/okta/okta-sdk-golang/okta/query")
      }
      imports.push("fmt");
      if (method.operation.responseModel !== undefined) {
        imports.push("fmt");
      }

      if (method.operation.bodyModel !== undefined) {
        imports.push("fmt")
      }
    }
  }

  if (object.model.crud !== undefined) {
    for (let method of object.model.crud) {
      if(method.operation.queryParams.length) {
        imports.push("github.com/okta/okta-sdk-golang/okta/query")
      }
      imports.push("fmt");
      if (method.operation.responseModel !== undefined) {
        imports.push("fmt");
      }

      if (method.operation.bodyModel !== undefined) {
        imports.push("fmt")
      }
    }
  }

  if (object.model.modelName === "LogEvent") {
    imports.push("github.com/okta/okta-sdk-golang/okta/query")
    imports.push("fmt")
  }

  imports = [...new Set(imports)];

  return imports;
}

function operationArgumentBuilder(operation) {
  const args = [];

  operation.pathParams.map((arg) => args.push(arg.name + " " + arg.type));

  if ((operation.method === 'post' || operation.method === 'put') && operation.bodyModel) {
    let bodyModel = ucFirst(_.camelCase(operation.bodyModel));

    if(bodyModel === "Application") {
      bodyModel = "App";
    }

    if(bodyModel === "Factor") {
      bodyModel = "UserFactor";
    }

    args.push(`body ` + bodyModel);
  }

  if(operation.operationId === "getApplication") {
    args.push(`appInstance App`);
  }

  if(operation.operationId === "getFactor") {
    args.push(`factorInstance UserFactor`);
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
  if ( operation.responseModel !== undefined ) {
    let responseModel = "*" +operation.responseModel
    if ( responseModel === "*Application" ) {
      if ( operation.operationId === "listApplications") {
        responseModel = "App"
      } else {
        responseModel = "interface{}"
      }
    }
    if ( responseModel === "*Factor" ) {
      if ( operation.operationId === "listFactors") {
        responseModel = "UserFactor"
      } else {
        responseModel = "interface{}"
      }
    }
    if ( operation.isArray !== undefined && operation.isArray === true) {
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
 return operationId ===  "listApplications" ||
    operationId === "listSupportedFactors" ||
    operationId === "listFactors";

}

function getClientTagResources(operations) {
  let tags = getClientTags(operations);
  let tagResources = []
  for (let tag of tags) {
    if (tag === "UserFactor") tag = "Factor";
    if (tag === "Log") tag = "LogEvent";
    tagResources.push(structProp(tag) + " *" + structProp(tag) + "Resource")
  }
  return tagResources.join("\n\t");
}

function getNewClientTagProps(operations) {
  let tags = getClientTags(operations);
  let tagResources = []
  for (let tag of tags) {
    if (tag === "UserFactor") tag = "Factor";
    if (tag === "Log") tag = "LogEvent";
    tagResources.push("c." + structProp(tag) + " = (*" + structProp(tag) + "Resource)(&c.resource)")
  }
  return tagResources.join("\n\t");
}

function buildModelProperties(model) {
  const properties = {};
  const finalProps = [];

  if(model.parent !== undefined) {
    for (let parentProperty of model.parent.properties) {
      properties[parentProperty.propertyName] = parentProperty;
    }
    if(model.parent.parent !== undefined) {
      for (let parentProperty of model.parent.parent.properties) {
        properties[parentProperty.propertyName] = parentProperty;
      }
    }
  }

  if(model.properties !== undefined) {
    for (let modelProperty of model.properties) {
      properties[modelProperty.propertyName] = modelProperty;
    }
  }

  for (let propKey in properties) {
    finalProps.push( structProp(properties[propKey].propertyName) + " " + getType(properties[propKey], "*") + " `json:\""+properties[propKey].propertyName+",omitempty\"`" );

  }

  return finalProps.join("\n\t");
}

function log(item) {
    console.log(item);
}



golang.process = ({ spec, operations, models, handlebars }) => {
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

  templates.push({
    src: 'templates/okta.go.hbs',
    dest: 'okta/okta.go',
    context: {
      "operations": operations,
      "models": models
    }
  });

  for (let model of models) {

    if(model.extends !== undefined) {
      model.parent = modelsByName[model.extends];

      if(model.parent.resolutionStrategy !== undefined && model.parent.parent == undefined) {
        for (let value in model.parent.resolutionStrategy.valueToModelMapping) {
          if (model.parent.resolutionStrategy.valueToModelMapping[value] === model.modelName) {
            model.resolution = {fieldName: model.parent.resolutionStrategy.propertyName, fieldValue: value};
          }
        }
      }

      if(model.parent.parent !== undefined && model.parent.parent.resolutionStrategy !== undefined) {
        model.resolution = model.parent.resolution;
      }
    }


    let modelOperations = {}

    if(model.crud != undefined) {
      for (let modelCrud of model.crud) {
        modelOperations[modelCrud.operation.operationId] = modelCrud.operation;
      }
    }

    for (let operation of operations) {
      let tag = operation.tags[0];
      if (tag === "UserFactor" ) tag = "Factor";
      if (tag === "Log" ) tag = "LogEvent";
      if (tag == model.modelName) {
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
    responseModelInterface
  });

  handlebars.registerPartial('partials.copyHeader', fs.readFileSync('generator/templates/partials/copyHeader.hbs', 'utf8'));
  handlebars.registerPartial('struct.withProp', fs.readFileSync('generator/templates/struct/withProp.go.hbs', 'utf8'));
  handlebars.registerPartial('model.imports', fs.readFileSync('generator/templates/model/imports.go.hbs', 'utf8'));
  handlebars.registerPartial('model.defaultMethod', fs.readFileSync('generator/templates/model/defaultMethod.go.hbs', 'utf8'));
  handlebars.registerPartial('model.defaultOperation', fs.readFileSync('generator/templates/model/defaultOperation.go.hbs', 'utf8'));

  fs.writeFile("generator/createdFiles.json", JSON.stringify(templates), function(error) {
    console.log(error);
  });
  return templates;
};
