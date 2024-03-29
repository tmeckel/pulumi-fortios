// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a list of `fortios.firewall.Addrgrp6`.
 */
export function getAddrgrp6list(args?: GetAddrgrp6listArgs, opts?: pulumi.InvokeOptions): Promise<GetAddrgrp6listResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:firewall/getAddrgrp6list:getAddrgrp6list", {
        "filter": args.filter,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getAddrgrp6list.
 */
export interface GetAddrgrp6listArgs {
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
 * A collection of values returned by getAddrgrp6list.
 */
export interface GetAddrgrp6listResult {
    readonly filter?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the `fortios.firewall.Addrgrp6`.
     */
    readonly namelists: string[];
    readonly vdomparam?: string;
}
/**
 * Provides a list of `fortios.firewall.Addrgrp6`.
 */
export function getAddrgrp6listOutput(args?: GetAddrgrp6listOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAddrgrp6listResult> {
    return pulumi.output(args).apply((a: any) => getAddrgrp6list(a, opts))
}

/**
 * A collection of arguments for invoking getAddrgrp6list.
 */
export interface GetAddrgrp6listOutputArgs {
    /**
     * A filter used to scope the list. See Filter results of datasource.
     */
    filter?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
