// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios router prefixlist6
 */
export function getPrefixlist6(args: GetPrefixlist6Args, opts?: pulumi.InvokeOptions): Promise<GetPrefixlist6Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:router/getPrefixlist6:getPrefixlist6", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getPrefixlist6.
 */
export interface GetPrefixlist6Args {
    /**
     * Specify the name of the desired router prefixlist6.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getPrefixlist6.
 */
export interface GetPrefixlist6Result {
    /**
     * Comment.
     */
    readonly comments: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name.
     */
    readonly name: string;
    /**
     * IPv6 prefix list rule. The structure of `rule` block is documented below.
     */
    readonly rules: outputs.router.GetPrefixlist6Rule[];
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios router prefixlist6
 */
export function getPrefixlist6Output(args: GetPrefixlist6OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPrefixlist6Result> {
    return pulumi.output(args).apply((a: any) => getPrefixlist6(a, opts))
}

/**
 * A collection of arguments for invoking getPrefixlist6.
 */
export interface GetPrefixlist6OutputArgs {
    /**
     * Specify the name of the desired router prefixlist6.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
