// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { InterfacePortArgs, InterfacePortState } from "./interfacePort";
export type InterfacePort = import("./interfacePort").InterfacePort;
export const InterfacePort: typeof import("./interfacePort").InterfacePort = null as any;
utilities.lazyLoad(exports, ["InterfacePort"], () => require("./interfacePort"));

export { RouteStaticArgs, RouteStaticState } from "./routeStatic";
export type RouteStatic = import("./routeStatic").RouteStatic;
export const RouteStatic: typeof import("./routeStatic").RouteStatic = null as any;
utilities.lazyLoad(exports, ["RouteStatic"], () => require("./routeStatic"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:networking/interfacePort:InterfacePort":
                return new InterfacePort(name, <any>undefined, { urn })
            case "fortios:networking/routeStatic:RouteStatic":
                return new RouteStatic(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "networking/interfacePort", _module)
pulumi.runtime.registerResourceModule("fortios", "networking/routeStatic", _module)
