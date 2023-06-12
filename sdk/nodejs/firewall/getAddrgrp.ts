// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios firewall addrgrp
 */
export function getAddrgrp(args: GetAddrgrpArgs, opts?: pulumi.InvokeOptions): Promise<GetAddrgrpResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:firewall/getAddrgrp:getAddrgrp", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getAddrgrp.
 */
export interface GetAddrgrpArgs {
    /**
     * Specify the name of the desired firewall addrgrp.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getAddrgrp.
 */
export interface GetAddrgrpResult {
    /**
     * Enable/disable use of this group in the static route configuration.
     */
    readonly allowRouting: string;
    /**
     * Tag category.
     */
    readonly category: string;
    /**
     * Color of icon on the GUI.
     */
    readonly color: number;
    /**
     * Comment.
     */
    readonly comment: string;
    /**
     * Enable/disable address exclusion.
     */
    readonly exclude: string;
    /**
     * Address exclusion member. The structure of `excludeMember` block is documented below.
     */
    readonly excludeMembers: outputs.firewall.GetAddrgrpExcludeMember[];
    /**
     * Security Fabric global object setting.
     */
    readonly fabricObject: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Address objects contained within the group. The structure of `member` block is documented below.
     */
    readonly members: outputs.firewall.GetAddrgrpMember[];
    /**
     * Tag name.
     */
    readonly name: string;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    readonly taggings: outputs.firewall.GetAddrgrpTagging[];
    /**
     * Address group type.
     */
    readonly type: string;
    /**
     * Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
     */
    readonly uuid: string;
    readonly vdomparam?: string;
    /**
     * Enable/disable address visibility in the GUI.
     */
    readonly visibility: string;
}
/**
 * Use this data source to get information on an fortios firewall addrgrp
 */
export function getAddrgrpOutput(args: GetAddrgrpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAddrgrpResult> {
    return pulumi.output(args).apply((a: any) => getAddrgrp(a, opts))
}

/**
 * A collection of arguments for invoking getAddrgrp.
 */
export interface GetAddrgrpOutputArgs {
    /**
     * Specify the name of the desired firewall addrgrp.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
