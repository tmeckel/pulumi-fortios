// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a list of `fortios.firewall.Policy64`.
 */
export function getPolicy64list(args?: GetPolicy64listArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicy64listResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:firewall/getPolicy64list:getPolicy64list", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicy64list.
 */
export interface GetPolicy64listArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getPolicy64list.
 */
export interface GetPolicy64listResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.firewall.Policy64`.
     */
    readonly policyidlists: number[];
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.firewall.Policy64`.
 */
export function getPolicy64listOutput(args?: GetPolicy64listOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicy64listResult> {
    return pulumi.output(args).apply((a: any) => getPolicy64list(a, opts))
}

/**
 * A collection of arguments for invoking getPolicy64list.
 */
export interface GetPolicy64listOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
