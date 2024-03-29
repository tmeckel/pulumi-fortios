// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure MS Exchange server entries. Applies to FortiOS Version `>= 6.2.4`.
 *
 * ## Import
 *
 * User Exchange can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:user/exchange:Exchange labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:user/exchange:Exchange labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Exchange extends pulumi.CustomResource {
    /**
     * Get an existing Exchange resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExchangeState, opts?: pulumi.CustomResourceOptions): Exchange {
        return new Exchange(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/exchange:Exchange';

    /**
     * Returns true if the given object is an instance of Exchange.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Exchange {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Exchange.__pulumiType;
    }

    /**
     * Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
     */
    public readonly authLevel!: pulumi.Output<string>;
    /**
     * Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
     */
    public readonly authType!: pulumi.Output<string>;
    /**
     * Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
     */
    public readonly autoDiscoverKdc!: pulumi.Output<string>;
    /**
     * Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
     */
    public readonly connectProtocol!: pulumi.Output<string>;
    /**
     * MS Exchange server fully qualified domain name.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
     */
    public readonly httpAuthType!: pulumi.Output<string>;
    /**
     * Server IPv4 address.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
     */
    public readonly kdcIps!: pulumi.Output<outputs.user.ExchangeKdcIp[] | undefined>;
    /**
     * MS Exchange server entry name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password for the specified username.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * MS Exchange server hostname.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    public readonly sslMinProtoVersion!: pulumi.Output<string>;
    /**
     * User name used to sign in to the server. Must have proper permissions for service.
     */
    public readonly username!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Exchange resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExchangeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExchangeArgs | ExchangeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExchangeState | undefined;
            resourceInputs["authLevel"] = state ? state.authLevel : undefined;
            resourceInputs["authType"] = state ? state.authType : undefined;
            resourceInputs["autoDiscoverKdc"] = state ? state.autoDiscoverKdc : undefined;
            resourceInputs["connectProtocol"] = state ? state.connectProtocol : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["httpAuthType"] = state ? state.httpAuthType : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["kdcIps"] = state ? state.kdcIps : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["serverName"] = state ? state.serverName : undefined;
            resourceInputs["sslMinProtoVersion"] = state ? state.sslMinProtoVersion : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as ExchangeArgs | undefined;
            resourceInputs["authLevel"] = args ? args.authLevel : undefined;
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["autoDiscoverKdc"] = args ? args.autoDiscoverKdc : undefined;
            resourceInputs["connectProtocol"] = args ? args.connectProtocol : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["httpAuthType"] = args ? args.httpAuthType : undefined;
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["kdcIps"] = args ? args.kdcIps : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["sslMinProtoVersion"] = args ? args.sslMinProtoVersion : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Exchange.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Exchange resources.
 */
export interface ExchangeState {
    /**
     * Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
     */
    authLevel?: pulumi.Input<string>;
    /**
     * Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
     */
    authType?: pulumi.Input<string>;
    /**
     * Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
     */
    autoDiscoverKdc?: pulumi.Input<string>;
    /**
     * Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
     */
    connectProtocol?: pulumi.Input<string>;
    /**
     * MS Exchange server fully qualified domain name.
     */
    domainName?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
     */
    httpAuthType?: pulumi.Input<string>;
    /**
     * Server IPv4 address.
     */
    ip?: pulumi.Input<string>;
    /**
     * KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
     */
    kdcIps?: pulumi.Input<pulumi.Input<inputs.user.ExchangeKdcIp>[]>;
    /**
     * MS Exchange server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for the specified username.
     */
    password?: pulumi.Input<string>;
    /**
     * MS Exchange server hostname.
     */
    serverName?: pulumi.Input<string>;
    /**
     * Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * User name used to sign in to the server. Must have proper permissions for service.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Exchange resource.
 */
export interface ExchangeArgs {
    /**
     * Authentication security level used for the RPC protocol layer. Valid values: `connect`, `call`, `packet`, `integrity`, `privacy`.
     */
    authLevel?: pulumi.Input<string>;
    /**
     * Authentication security type used for the RPC protocol layer. Valid values: `spnego`, `ntlm`, `kerberos`.
     */
    authType?: pulumi.Input<string>;
    /**
     * Enable/disable automatic discovery of KDC IP addresses. Valid values: `enable`, `disable`.
     */
    autoDiscoverKdc?: pulumi.Input<string>;
    /**
     * Connection protocol used to connect to MS Exchange service. Valid values: `rpc-over-tcp`, `rpc-over-http`, `rpc-over-https`.
     */
    connectProtocol?: pulumi.Input<string>;
    /**
     * MS Exchange server fully qualified domain name.
     */
    domainName?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Authentication security type used for the HTTP transport. Valid values: `basic`, `ntlm`.
     */
    httpAuthType?: pulumi.Input<string>;
    /**
     * Server IPv4 address.
     */
    ip?: pulumi.Input<string>;
    /**
     * KDC IPv4 addresses for Kerberos authentication. The structure of `kdcIp` block is documented below.
     */
    kdcIps?: pulumi.Input<pulumi.Input<inputs.user.ExchangeKdcIp>[]>;
    /**
     * MS Exchange server entry name.
     */
    name?: pulumi.Input<string>;
    /**
     * Password for the specified username.
     */
    password?: pulumi.Input<string>;
    /**
     * MS Exchange server hostname.
     */
    serverName?: pulumi.Input<string>;
    /**
     * Minimum SSL/TLS protocol version for HTTPS transport (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
     */
    sslMinProtoVersion?: pulumi.Input<string>;
    /**
     * User name used to sign in to the server. Must have proper permissions for service.
     */
    username?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
