const purify = require("purify-css");

const htmlFiles = ["./build/*.html", "./build/static/js/*.js"];

const cssFiles = ["./build/static/css/*.css"];

const opts = {
  output: "purified.css",
  //   minify: true,
};

purify(htmlFiles, cssFiles, opts, function (res) {
  log(res);
});
