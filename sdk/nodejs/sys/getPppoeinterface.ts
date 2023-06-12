// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system pppoeinterface
 */
export function getPppoeinterface(args: GetPppoeinterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetPppoeinterfaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:sys/getPppoeinterface:getPppoeinterface", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getPppoeinterface.
 */
export interface GetPppoeinterfaceArgs {
    /**
     * Specify the name of the desired system pppoeinterface.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getPppoeinterface.
 */
export interface GetPppoeinterfaceResult {
    /**
     * PPPoE AC name.
     */
    readonly acName: string;
    /**
     * PPP authentication type to use.
     */
    readonly authType: string;
    /**
     * Name for the physical interface.
     */
    readonly device: string;
    /**
     * Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface.
     */
    readonly dialOnDemand: string;
    /**
     * PPPoE discovery init timeout value in (0-4294967295 sec).
     */
    readonly discRetryTimeout: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * PPPoE auto disconnect after idle timeout (0-4294967295 sec).
     */
    readonly idleTimeout: number;
    /**
     * PPPoE unnumbered IP.
     */
    readonly ipunnumbered: string;
    /**
     * Enable/disable IPv6 Control Protocol (IPv6CP).
     */
    readonly ipv6: string;
    /**
     * PPPoE LCP echo interval in (0-4294967295 sec, default = 5).
     */
    readonly lcpEchoInterval: number;
    /**
     * Maximum missed LCP echo messages before disconnect (0-4294967295, default = 3).
     */
    readonly lcpMaxEchoFails: number;
    /**
     * Name of the PPPoE interface.
     */
    readonly name: string;
    /**
     * PPPoE terminate timeout value in (0-4294967295 sec).
     */
    readonly padtRetryTimeout: number;
    /**
     * Enter the password.
     */
    readonly password: string;
    /**
     * Enable/disable PPPoE unnumbered negotiation.
     */
    readonly pppoeUnnumberedNegotiate: string;
    /**
     * PPPoE service name.
     */
    readonly serviceName: string;
    /**
     * User name.
     */
    readonly username: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system pppoeinterface
 */
export function getPppoeinterfaceOutput(args: GetPppoeinterfaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPppoeinterfaceResult> {
    return pulumi.output(args).apply((a: any) => getPppoeinterface(a, opts))
}

/**
 * A collection of arguments for invoking getPppoeinterface.
 */
export interface GetPppoeinterfaceOutputArgs {
    /**
     * Specify the name of the desired system pppoeinterface.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}