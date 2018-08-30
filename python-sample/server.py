import os
import platform
import json

#Getting App Environment Info
from cfenv import AppEnv
from flask import render_template
from flask import Flask
app = Flask(__name__)

# Get port from environment variable or choose 9099 as local default
port = int(os.getenv('VCAP_APP_PORT', 8080))
@app.route('/')
def hello():
    env = AppEnv()
    app_name = env.name
    app_uris = env.uris
    space_name = env.space
    index = env.index
    vcap_app_env = json.loads(os.getenv('VCAP_APPLICATION', '{}'))
    app_mem_limit = str(vcap_app_env["limits"].get('mem'))
    app_disk_limit = str(vcap_app_env["limits"].get('disk'))
    #return render_template('index1.html', app_name=app_name, app_uris=app_uris)
    return render_template('index.html', app_name=app_name, app_uris=app_uris, space_name=space_name, index=index, app_mem_limit=app_mem_limit, app_disk_limit=app_disk_limit)
    #return app.send_static_file('index.html')
    #return 'Appname = ' + str(app_name) + ' URIs = ' + str(app_uris) + ' space = ' + str(space_name) + ' instance index = ' + str (index) + ' mem limit = ' + app_mem_limit + ' disk limit = ' + app_disk_limit

if __name__ == '__main__':
        # Run the app, listening on all IPs with our chosen port number
    app.run(host='0.0.0.0', port=port)
