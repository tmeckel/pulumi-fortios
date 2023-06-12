// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { FilterArgs, FilterState } from "./filter";
export type Filter = import("./filter").Filter;
export const Filter: typeof import("./filter").Filter = null as any;
utilities.lazyLoad(exports, ["Filter"], () => require("./filter"));

export { OverridefilterArgs, OverridefilterState } from "./overridefilter";
export type Overridefilter = import("./overridefilter").Overridefilter;
export const Overridefilter: typeof import("./overridefilter").Overridefilter = null as any;
utilities.lazyLoad(exports, ["Overridefilter"], () => require("./overridefilter"));

export { OverridesettingArgs, OverridesettingState } from "./overridesetting";
export type Overridesetting = import("./overridesetting").Overridesetting;
export const Overridesetting: typeof import("./overridesetting").Overridesetting = null as any;
utilities.lazyLoad(exports, ["Overridesetting"], () => require("./overridesetting"));

export { SettingArgs, SettingState } from "./setting";
export type Setting = import("./setting").Setting;
export const Setting: typeof import("./setting").Setting = null as any;
utilities.lazyLoad(exports, ["Setting"], () => require("./setting"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:logfortianalyzer/filter:Filter":
                return new Filter(name, <any>undefined, { urn })
            case "fortios:logfortianalyzer/overridefilter:Overridefilter":
                return new Overridefilter(name, <any>undefined, { urn })
            case "fortios:logfortianalyzer/overridesetting:Overridesetting":
                return new Overridesetting(name, <any>undefined, { urn })
            case "fortios:logfortianalyzer/setting:Setting":
                return new Setting(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "logfortianalyzer/filter", _module)
pulumi.runtime.registerResourceModule("fortios", "logfortianalyzer/overridefilter", _module)
pulumi.runtime.registerResourceModule("fortios", "logfortianalyzer/overridesetting", _module)
pulumi.runtime.registerResourceModule("fortios", "logfortianalyzer/setting", _module)