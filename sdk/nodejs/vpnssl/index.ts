// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ClientArgs, ClientState } from "./client";
export type Client = import("./client").Client;
export const Client: typeof import("./client").Client = null as any;
utilities.lazyLoad(exports, ["Client"], () => require("./client"));

export { SettingsArgs, SettingsState } from "./settings";
export type Settings = import("./settings").Settings;
export const Settings: typeof import("./settings").Settings = null as any;
utilities.lazyLoad(exports, ["Settings"], () => require("./settings"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:vpnssl/client:Client":
                return new Client(name, <any>undefined, { urn })
            case "fortios:vpnssl/settings:Settings":
                return new Settings(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "vpnssl/client", _module)
pulumi.runtime.registerResourceModule("fortios", "vpnssl/settings", _module)