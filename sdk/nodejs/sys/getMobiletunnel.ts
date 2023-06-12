// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system mobiletunnel
 */
export function getMobiletunnel(args: GetMobiletunnelArgs, opts?: pulumi.InvokeOptions): Promise<GetMobiletunnelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:sys/getMobiletunnel:getMobiletunnel", {
        "name": args.name,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getMobiletunnel.
 */
export interface GetMobiletunnelArgs {
    /**
     * Specify the name of the desired system mobiletunnel.
     */
    name: string;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getMobiletunnel.
 */
export interface GetMobiletunnelResult {
    /**
     * Hash Algorithm (Keyed MD5).
     */
    readonly hashAlgorithm: string;
    /**
     * Home IP address (Format: xxx.xxx.xxx.xxx).
     */
    readonly homeAddress: string;
    /**
     * IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
     */
    readonly homeAgent: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
     */
    readonly lifetime: number;
    /**
     * NEMO authentication key.
     */
    readonly nMhaeKey: string;
    /**
     * NEMO authentication key type (ascii or base64).
     */
    readonly nMhaeKeyType: string;
    /**
     * NEMO authentication SPI (default: 256).
     */
    readonly nMhaeSpi: number;
    /**
     * Tunnel name.
     */
    readonly name: string;
    /**
     * NEMO network configuration. The structure of `network` block is documented below.
     */
    readonly networks: outputs.sys.GetMobiletunnelNetwork[];
    /**
     * NMMO HA registration interval (5 - 300, default = 5).
     */
    readonly regInterval: number;
    /**
     * Maximum number of NMMO HA registration retries (1 to 30, default = 3).
     */
    readonly regRetry: number;
    /**
     * Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
     */
    readonly renewInterval: number;
    /**
     * Select the associated interface name from available options.
     */
    readonly roamingInterface: string;
    /**
     * Enable/disable this mobile tunnel.
     */
    readonly status: string;
    /**
     * NEMO tunnnel mode (GRE tunnel).
     */
    readonly tunnelMode: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system mobiletunnel
 */
export function getMobiletunnelOutput(args: GetMobiletunnelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMobiletunnelResult> {
    return pulumi.output(args).apply((a: any) => getMobiletunnel(a, opts))
}

/**
 * A collection of arguments for invoking getMobiletunnel.
 */
export interface GetMobiletunnelOutputArgs {
    /**
     * Specify the name of the desired system mobiletunnel.
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
