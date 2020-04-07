let fs = require("fs")

let filesJson = fs.readFileSync("./generator/createdFiles.json")

try {
  for (let file of JSON.parse(filesJson)) {
    fs.unlinkSync("../" + file.dest)
  }
} catch (e) {

}

fs.writeFile("./generator/createdFiles.json", JSON.stringify(""), function(error) {
  console.log(error)
});
