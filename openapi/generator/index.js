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

function getType(obj) {
  switch(obj.commonType) {
    case 'dateTime' :
      return String.raw`time.Time`;
    case 'integer' :
      return String.raw`int64`;
    case 'boolean' :
      return String.raw`bool`;
    case 'hash' :
    case 'array' :
      return String.raw`map[string]interface{}`;
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

function structProp(prop) {
  prop = prop.replace(/#/g,"");
  prop = prop.replace(/(\_\w)/g, function(m){return m[1].toUpperCase();});

  prop = prop.charAt(0).toUpperCase() + prop.slice(1);

  return prop;
}

function getImports(operations, models, tag) {
  let imports = [];

  imports.push("fmt");
  imports.push("encoding/json");

  for (let model of models) {
    if (model.tags[0] == tag) {
      if (model.properties !== undefined) {
        for (let property of model.properties) {
          switch (property.commonType) {
            case 'dateTime' :
              imports.push("time");
          }
        }
      }
    }
  }

  for (let operation of operations) {
    if (operation.tags[0] == tag) {
      if (operation.responseModel !== undefined) {
        imports.push("fmt");
        imports.push("encoding/json");
      }

      if (operation.bodyModel !== undefined) {
        if(tag == "Group") {
          // console.log(operation.operationId, operation.bodyModel);
        }
        imports.push("bytes")
      }
    }
  }

  return [...new Set(imports)];
}

function operationArgumentBuilder(operation) {
  const args = [];

  operation.pathParams.map((arg) => args.push(arg.name + " " + arg.type));

  if ((operation.method === 'post' || operation.method === 'put') && operation.bodyModel) {
    args.push(`body ` + ucFirst(_.camelCase(operation.bodyModel)));
  }

  // if (operation.queryParams.length) {
  //   args.push('queryParameters');
  // }

  return args.join(', ');
}

function getPath(operation) {
  let path = operation.path;
  for (let param of operation.pathParams) {
    path = path.replace(`{${param.name}}`, String.raw`"+${param.name}+"`);
  }

  return `${path}`;
}

function returnType(operation) {
  if ( operation.responseModel !== undefined ) {
    return " *" + operation.responseModel + " ";
  }
  return " ";
}

function log(item) {
  console.log(item);
}



golang.process = ({ spec, operations, models, handlebars }) => {

  golang.spec = spec;
  const templates = [];

  for (let model of models) {
    if(model.modelName == "Application") {
      // console.log(model);
    }
    templates.push({
      src: 'templates/model.go.hbs',
      dest: 'okta/' + lowercaseFirstLetter(model.modelName) + '.go',
      context: {
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
    returnType
  });

  handlebars.registerPartial('partials.copyHeader', fs.readFileSync('generator/templates/partials/copyHeader.hbs', 'utf8'));
  handlebars.registerPartial('struct.withProp', fs.readFileSync('generator/templates/struct/withProp.go.hbs', 'utf8'));
  handlebars.registerPartial('model.imports', fs.readFileSync('generator/templates/model/imports.go.hbs', 'utf8'));

  fs.writeFile("generator/createdFiles.json", JSON.stringify(templates), function(error) {
    console.log(error);
  });
  return templates;
};
