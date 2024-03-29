// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an fortios system ddns
 */
export function getDdns(args: GetDdnsArgs, opts?: pulumi.InvokeOptions): Promise<GetDdnsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:sys/getDdns:getDdns", {
        "ddnsid": args.ddnsid,
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getDdns.
 */
export interface GetDdnsArgs {
    /**
     * Specify the ddnsid of the desired system ddns.
     */
    ddnsid: number;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getDdns.
 */
export interface GetDdnsResult {
    /**
     * Address type of interface address in DDNS update.
     */
    readonly addrType: string;
    /**
     * Bound IP address.
     */
    readonly boundIp: string;
    /**
     * Enable/disable use of clear text connections.
     */
    readonly clearText: string;
    /**
     * Enable/disable TSIG authentication for your DDNS server.
     */
    readonly ddnsAuth: string;
    /**
     * Your fully qualified domain name (for example, yourname.DDNS.com).
     */
    readonly ddnsDomain: string;
    /**
     * DDNS update key (base 64 encoding).
     */
    readonly ddnsKey: string;
    /**
     * DDNS update key name.
     */
    readonly ddnsKeyname: string;
    /**
     * DDNS password.
     */
    readonly ddnsPassword: string;
    /**
     * Select a DDNS service provider.
     */
    readonly ddnsServer: string;
    /**
     * Generic DDNS server IP/FQDN list. The structure of `ddnsServerAddr` block is documented below.
     */
    readonly ddnsServerAddrs: outputs.sys.GetDdnsDdnsServerAddr[];
    /**
     * Generic DDNS server IP.
     */
    readonly ddnsServerIp: string;
    /**
     * DDNS Serial Number.
     */
    readonly ddnsSn: string;
    /**
     * Time-to-live for DDNS packets.
     */
    readonly ddnsTtl: number;
    /**
     * DDNS user name.
     */
    readonly ddnsUsername: string;
    /**
     * Zone of your domain name (for example, DDNS.com).
     */
    readonly ddnsZone: string;
    /**
     * DDNS ID.
     */
    readonly ddnsid: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Monitored interface. The structure of `monitorInterface` block is documented below.
     */
    readonly monitorInterfaces: outputs.sys.GetDdnsMonitorInterface[];
    /**
     * Address type of the DDNS server.
     */
    readonly serverType: string;
    /**
     * Name of local certificate for SSL connections.
     */
    readonly sslCertificate: string;
    /**
     * DDNS update interval (60 - 2592000 sec, default = 300).
     */
    readonly updateInterval: number;
    /**
     * Enable/disable use of public IP address.
     */
    readonly usePublicIp: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on an fortios system ddns
 */
export function getDdnsOutput(args: GetDdnsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDdnsResult> {
    return pulumi.output(args).apply((a: any) => getDdns(a, opts))
}

/**
 * A collection of arguments for invoking getDdns.
 */
export interface GetDdnsOutputArgs {
    /**
     * Specify the ddnsid of the desired system ddns.
     */
    ddnsid: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
