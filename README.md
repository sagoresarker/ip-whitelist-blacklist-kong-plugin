## Issues (fixed in the previous commit)

In this project i implemented a kong plugin which controll the ip whitelisting, blacklisting. But I faced a issue repeatedly and could not manage to solve it properly.

The issue i faced during the r&d was,

```sh
kong-migrations-1 exited with code 0
kong-1             | 2025/02/16 04:02:38 [warn] 1#0: the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /usr/local/kong/nginx.conf:7
kong-1             | nginx: [warn] the "user" directive makes sense only if the master process runs with super-user privileges, ignored in /usr/local/kong/nginx.conf:7
kong-1             | 2025/02/16 04:02:38 [notice] 1#0: [kong] process.lua:90 [pluginserver] loading external plugins info
kong-1             | sh: 1: /usr/local/kong/plugins/ip-tacker/query: not found
kong-1             | 2025/02/16 04:02:38 [error] 1#0: init_by_lua error: ...local/share/lua/5.1/kong/runloop/plugin_servers/init.lua:37: failed loading external plugins: failed decoding plugin info: Expected value but found T_END at character 1
kong-1             | stack traceback:
kong-1             |    [C]: in function 'assert'
kong-1             |    ...local/share/lua/5.1/kong/runloop/plugin_servers/init.lua:37: in function 'get_plugin'
kong-1             |    ...local/share/lua/5.1/kong/runloop/plugin_servers/init.lua:43: in function 'load_plugin'
kong-1             |    /usr/local/share/lua/5.1/kong/db/dao/plugins.lua:172: in function 'load_plugin_handler'
kong-1             |    /usr/local/share/lua/5.1/kong/db/dao/plugins.lua:276: in function 'load_plugin'
kong-1             |    /usr/local/share/lua/5.1/kong/db/dao/plugins.lua:328: in function 'load_plugin_schemas'
kong-1             |    /usr/local/share/lua/5.1/kong/init.lua:708: in function 'init'
kong-1             |    init_by_lua(nginx-kong.conf:44):3: in main chunk
kong-1             | nginx: [error] init_by_lua error: ...local/share/lua/5.1/kong/runloop/plugin_servers/init.lua:37: failed loading external plugins: failed decoding plugin info: Expected value but found T_END at character 1
kong-1             | stack traceback:
kong-1             |    [C]: in function 'assert'
kong-1             |    ...local/share/lua/5.1/kong/runloop/plugin_servers/init.lua:37: in function 'get_plugin'
kong-1             |    ...local/share/lua/5.1/kong/runloop/plugin_servers/init.lua:43: in function 'load_plugin'
kong-1             |    /usr/local/share/lua/5.1/kong/db/dao/plugins.lua:172: in function 'load_plugin_handler'
kong-1             |    /usr/local/share/lua/5.1/kong/db/dao/plugins.lua:276: in function 'load_plugin'
kong-1             |    /usr/local/share/lua/5.1/kong/db/dao/plugins.lua:328: in function 'load_plugin_schemas'
kong-1             |    /usr/local/share/lua/5.1/kong/init.lua:708: in function 'init'
kong-1             |    init_by_lua(nginx-kong.conf:44):3: in main chunk
kong-1 exited with code 1

```

## How to run this application
```sh
docker compose up
```