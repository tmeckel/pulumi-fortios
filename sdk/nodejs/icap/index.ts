// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ProfileArgs, ProfileState } from "./profile";
export type Profile = import("./profile").Profile;
export const Profile: typeof import("./profile").Profile = null as any;
utilities.lazyLoad(exports, ["Profile"], () => require("./profile"));

export { ServerArgs, ServerState } from "./server";
export type Server = import("./server").Server;
export const Server: typeof import("./server").Server = null as any;
utilities.lazyLoad(exports, ["Server"], () => require("./server"));

export { ServergroupArgs, ServergroupState } from "./servergroup";
export type Servergroup = import("./servergroup").Servergroup;
export const Servergroup: typeof import("./servergroup").Servergroup = null as any;
utilities.lazyLoad(exports, ["Servergroup"], () => require("./servergroup"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:icap/profile:Profile":
                return new Profile(name, <any>undefined, { urn })
            case "fortios:icap/server:Server":
                return new Server(name, <any>undefined, { urn })
            case "fortios:icap/servergroup:Servergroup":
                return new Servergroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "icap/profile", _module)
pulumi.runtime.registerResourceModule("fortios", "icap/server", _module)
pulumi.runtime.registerResourceModule("fortios", "icap/servergroup", _module)
