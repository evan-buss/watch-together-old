const express = require('express')
var compression = require('compression')
const app = express()
const port = process.env.PORT || 80


app.use(compression())
app.use(express.static('public'))

app.get("/", function(req, res) {
  res.sendFile("public/index.html")
})

//The 404 Route (ALWAYS Keep this as the last route)
app.get('*', function(req, res){
  res.redirect("/")
});

app.listen(port, () => console.log(`Example app listening on port ${port}!`))