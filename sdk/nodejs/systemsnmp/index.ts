// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CommunityArgs, CommunityState } from "./community";
export type Community = import("./community").Community;
export const Community: typeof import("./community").Community = null as any;
utilities.lazyLoad(exports, ["Community"], () => require("./community"));

export { GetCommunityArgs, GetCommunityResult, GetCommunityOutputArgs } from "./getCommunity";
export const getCommunity: typeof import("./getCommunity").getCommunity = null as any;
export const getCommunityOutput: typeof import("./getCommunity").getCommunityOutput = null as any;
utilities.lazyLoad(exports, ["getCommunity","getCommunityOutput"], () => require("./getCommunity"));

export { GetCommunitylistArgs, GetCommunitylistResult, GetCommunitylistOutputArgs } from "./getCommunitylist";
export const getCommunitylist: typeof import("./getCommunitylist").getCommunitylist = null as any;
export const getCommunitylistOutput: typeof import("./getCommunitylist").getCommunitylistOutput = null as any;
utilities.lazyLoad(exports, ["getCommunitylist","getCommunitylistOutput"], () => require("./getCommunitylist"));

export { GetSysinfoArgs, GetSysinfoResult, GetSysinfoOutputArgs } from "./getSysinfo";
export const getSysinfo: typeof import("./getSysinfo").getSysinfo = null as any;
export const getSysinfoOutput: typeof import("./getSysinfo").getSysinfoOutput = null as any;
utilities.lazyLoad(exports, ["getSysinfo","getSysinfoOutput"], () => require("./getSysinfo"));

export { GetUserArgs, GetUserResult, GetUserOutputArgs } from "./getUser";
export const getUser: typeof import("./getUser").getUser = null as any;
export const getUserOutput: typeof import("./getUser").getUserOutput = null as any;
utilities.lazyLoad(exports, ["getUser","getUserOutput"], () => require("./getUser"));

export { GetUserlistArgs, GetUserlistResult, GetUserlistOutputArgs } from "./getUserlist";
export const getUserlist: typeof import("./getUserlist").getUserlist = null as any;
export const getUserlistOutput: typeof import("./getUserlist").getUserlistOutput = null as any;
utilities.lazyLoad(exports, ["getUserlist","getUserlistOutput"], () => require("./getUserlist"));

export { MibviewArgs, MibviewState } from "./mibview";
export type Mibview = import("./mibview").Mibview;
export const Mibview: typeof import("./mibview").Mibview = null as any;
utilities.lazyLoad(exports, ["Mibview"], () => require("./mibview"));

export { SysinfoArgs, SysinfoState } from "./sysinfo";
export type Sysinfo = import("./sysinfo").Sysinfo;
export const Sysinfo: typeof import("./sysinfo").Sysinfo = null as any;
utilities.lazyLoad(exports, ["Sysinfo"], () => require("./sysinfo"));

export { UserArgs, UserState } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "fortios:systemsnmp/community:Community":
                return new Community(name, <any>undefined, { urn })
            case "fortios:systemsnmp/mibview:Mibview":
                return new Mibview(name, <any>undefined, { urn })
            case "fortios:systemsnmp/sysinfo:Sysinfo":
                return new Sysinfo(name, <any>undefined, { urn })
            case "fortios:systemsnmp/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("fortios", "systemsnmp/community", _module)
pulumi.runtime.registerResourceModule("fortios", "systemsnmp/mibview", _module)
pulumi.runtime.registerResourceModule("fortios", "systemsnmp/sysinfo", _module)
pulumi.runtime.registerResourceModule("fortios", "systemsnmp/user", _module)
