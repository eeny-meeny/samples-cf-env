var cfenv = require("cfenv");
var appEnv = cfenv.getAppEnv();
var vcap_application = appEnv.app;
var express = require( 'express');
var app = express();

app.use(express.static( __dirname + '/static'));
app.set('view engine', 'ejs');

app.get( '/', function ( req, res) {
  res.render(__dirname + "/static/index.ejs", {
    application_name:   appEnv.name,
    app_uris:           appEnv.url,
    app_space_name:     vcap_application.space_name,
    app_mem_limits:     vcap_application.limits.mem,
    app_disk_limits:    vcap_application.limits.disk
  });
});

app.listen( process.env.PORT || 4000);
