let fs = require("fs")

let filesJson = fs.readFileSync("./generator/createdFiles.json")

for ( let file of JSON.parse(filesJson) ) {
  fs.unlinkSync("../"+file.dest)
}

fs.writeFile("./generator/createdFiles.json", JSON.stringify(""), function(error) {
  console.log(error)
});
