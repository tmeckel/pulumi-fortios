// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on fortios system emailserver
 */
export function getEmailserver(args?: GetEmailserverArgs, opts?: pulumi.InvokeOptions): Promise<GetEmailserverResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("fortios:sys/getEmailserver:getEmailserver", {
        "vdomparam": args.vdomparam,
    }, opts);
}

/**
 * A collection of arguments for invoking getEmailserver.
 */
export interface GetEmailserverArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: string;
}

/**
 * A collection of values returned by getEmailserver.
 */
export interface GetEmailserverResult {
    /**
     * Enable/disable authentication.
     */
    readonly authenticate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Specify outgoing interface to reach server.
     */
    readonly interface: string;
    /**
     * Specify how to select outgoing interface to reach server.
     */
    readonly interfaceSelectMethod: string;
    /**
     * SMTP server user password for authentication.
     */
    readonly password: string;
    /**
     * SMTP server port.
     */
    readonly port: number;
    /**
     * Reply-To email address.
     */
    readonly replyTo: string;
    /**
     * Connection security used by the email server.
     */
    readonly security: string;
    /**
     * SMTP server IP address or hostname.
     */
    readonly server: string;
    /**
     * SMTP server IPv4 source IP.
     */
    readonly sourceIp: string;
    /**
     * SMTP server IPv6 source IP.
     */
    readonly sourceIp6: string;
    /**
     * Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
     */
    readonly sslMinProtoVersion: string;
    /**
     * Use FortiGuard Message service or custom email server.
     */
    readonly type: string;
    /**
     * SMTP server user name for authentication.
     */
    readonly username: string;
    /**
     * Enable/disable validation of server certificate.
     */
    readonly validateServer: string;
    readonly vdomparam?: string;
}
/**
 * Use this data source to get information on fortios system emailserver
 */
export function getEmailserverOutput(args?: GetEmailserverOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEmailserverResult> {
    return pulumi.output(args).apply((a: any) => getEmailserver(a, opts))
}

/**
 * A collection of arguments for invoking getEmailserver.
 */
export interface GetEmailserverOutputArgs {
    /**
     * Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
