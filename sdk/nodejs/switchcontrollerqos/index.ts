// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { Dot1pmapArgs, Dot1pmapState } from "./dot1pmap";
export type Dot1pmap = import("./dot1pmap").Dot1pmap;
export const Dot1pmap: typeof import("./dot1pmap").Dot1pmap = null as any;
utilities.lazyLoad(exports, ["Dot1pmap"], () => require("./dot1pmap"));

export { IpdscpmapArgs, IpdscpmapState } from "./ipdscpmap";
export type Ipdscpmap = import("./ipdscpmap").Ipdscpmap;
export const Ipdscpmap: typeof import("./ipdscpmap").Ipdscpmap = null as any;
utilities.lazyLoad(exports, ["Ipdscpmap"], () => require("./ipdscpmap"));

export { QospolicyArgs, QospolicyState } from "./qospolicy";
export type Qospolicy = import("./qospolicy").Qospolicy;
export const Qospolicy: typeof import("./qospolicy").Qospolicy = null as any;
utilities.lazyLoad(exports, ["Qospolicy"], () => require("./qospolicy"));

export { QueuepolicyArgs, QueuepolicyState } from "./queuepolicy";
export type Queuepolicy = import("./queuepolicy").Queuepolicy;
export const Queuepolicy: typeof import("./queuepolicy").Queuepolicy = null as any;
utilities.lazyLoad(exports, ["Queuepolicy"], () => require("./queuepolicy"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:switchcontrollerqos/dot1pmap:Dot1pmap":
                return new Dot1pmap(name, <any>undefined, { urn })
            case "fortios:switchcontrollerqos/ipdscpmap:Ipdscpmap":
                return new Ipdscpmap(name, <any>undefined, { urn })
            case "fortios:switchcontrollerqos/qospolicy:Qospolicy":
                return new Qospolicy(name, <any>undefined, { urn })
            case "fortios:switchcontrollerqos/queuepolicy:Queuepolicy":
                return new Queuepolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "switchcontrollerqos/dot1pmap", _module)
pulumi.runtime.registerResourceModule("fortios", "switchcontrollerqos/ipdscpmap", _module)
pulumi.runtime.registerResourceModule("fortios", "switchcontrollerqos/qospolicy", _module)
pulumi.runtime.registerResourceModule("fortios", "switchcontrollerqos/queuepolicy", _module)
