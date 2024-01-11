import "./textcoding.js";

import { protocGenConnectEs } from "@connectrpc/protoc-gen-connect-es/src/protoc-gen-connect-es-plugin.js";

import { runQuickJs  } from "./run-quickjs.js";

runQuickJs(protocGenConnectEs);
