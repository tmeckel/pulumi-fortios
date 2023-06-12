// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { FilterArgs, FilterState } from "./filter";
export type Filter = import("./filter").Filter;
export const Filter: typeof import("./filter").Filter = null as any;
utilities.lazyLoad(exports, ["Filter"], () => require("./filter"));

export { SettingArgs, SettingState } from "./setting";
export type Setting = import("./setting").Setting;
export const Setting: typeof import("./setting").Setting = null as any;
utilities.lazyLoad(exports, ["Setting"], () => require("./setting"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:logdisk/filter:Filter":
                return new Filter(name, <any>undefined, { urn })
            case "fortios:logdisk/setting:Setting":
                return new Setting(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "logdisk/filter", _module)
pulumi.runtime.registerResourceModule("fortios", "logdisk/setting", _module)
