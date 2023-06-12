// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * SNMP community configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.systemsnmp.Community("trname", {
 *     events: "cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-pattern av-fragmented fm-if-change bgp-established bgp-backward-transition ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-disconnect wc-ap-up wc-ap-down fswctl-session-up fswctl-session-down load-balance-real-server-down per-cpu-high",
 *     fosid: 1,
 *     queryV1Port: 161,
 *     queryV1Status: "enable",
 *     queryV2cPort: 161,
 *     queryV2cStatus: "enable",
 *     status: "enable",
 *     trapV1Lport: 162,
 *     trapV1Rport: 162,
 *     trapV1Status: "enable",
 *     trapV2cLport: 162,
 *     trapV2cRport: 162,
 *     trapV2cStatus: "enable",
 * });
 * ```
 *
 * ## Import
 *
 * SystemSnmp Community can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:systemsnmp/community:Community labelname {{fosid}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:systemsnmp/community:Community labelname {{fosid}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Community extends pulumi.CustomResource {
    /**
     * Get an existing Community resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CommunityState, opts?: pulumi.CustomResourceOptions): Community {
        return new Community(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:systemsnmp/community:Community';

    /**
     * Returns true if the given object is an instance of Community.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Community {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Community.__pulumiType;
    }

    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * SNMP trap events.
     */
    public readonly events!: pulumi.Output<string>;
    /**
     * Community ID.
     */
    public readonly fosid!: pulumi.Output<number>;
    /**
     * Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
     */
    public readonly hosts!: pulumi.Output<outputs.systemsnmp.CommunityHost[] | undefined>;
    /**
     * Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
     */
    public readonly hosts6s!: pulumi.Output<outputs.systemsnmp.CommunityHosts6[] | undefined>;
    /**
     * SNMP access control MIB view.
     */
    public readonly mibView!: pulumi.Output<string>;
    /**
     * Community name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * SNMP v1 query port (default = 161).
     */
    public readonly queryV1Port!: pulumi.Output<number>;
    /**
     * Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
     */
    public readonly queryV1Status!: pulumi.Output<string>;
    /**
     * SNMP v2c query port (default = 161).
     */
    public readonly queryV2cPort!: pulumi.Output<number>;
    /**
     * Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
     */
    public readonly queryV2cStatus!: pulumi.Output<string>;
    /**
     * Enable/disable this SNMP community. Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * SNMP v1 trap local port (default = 162).
     */
    public readonly trapV1Lport!: pulumi.Output<number>;
    /**
     * SNMP v1 trap remote port (default = 162).
     */
    public readonly trapV1Rport!: pulumi.Output<number>;
    /**
     * Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
     */
    public readonly trapV1Status!: pulumi.Output<string>;
    /**
     * SNMP v2c trap local port (default = 162).
     */
    public readonly trapV2cLport!: pulumi.Output<number>;
    /**
     * SNMP v2c trap remote port (default = 162).
     */
    public readonly trapV2cRport!: pulumi.Output<number>;
    /**
     * Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
     */
    public readonly trapV2cStatus!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * SNMP access control VDOMs. The structure of `vdoms` block is documented below.
     */
    public readonly vdoms!: pulumi.Output<outputs.systemsnmp.CommunityVdom[] | undefined>;

    /**
     * Create a Community resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CommunityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CommunityArgs | CommunityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CommunityState | undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["events"] = state ? state.events : undefined;
            resourceInputs["fosid"] = state ? state.fosid : undefined;
            resourceInputs["hosts"] = state ? state.hosts : undefined;
            resourceInputs["hosts6s"] = state ? state.hosts6s : undefined;
            resourceInputs["mibView"] = state ? state.mibView : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["queryV1Port"] = state ? state.queryV1Port : undefined;
            resourceInputs["queryV1Status"] = state ? state.queryV1Status : undefined;
            resourceInputs["queryV2cPort"] = state ? state.queryV2cPort : undefined;
            resourceInputs["queryV2cStatus"] = state ? state.queryV2cStatus : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["trapV1Lport"] = state ? state.trapV1Lport : undefined;
            resourceInputs["trapV1Rport"] = state ? state.trapV1Rport : undefined;
            resourceInputs["trapV1Status"] = state ? state.trapV1Status : undefined;
            resourceInputs["trapV2cLport"] = state ? state.trapV2cLport : undefined;
            resourceInputs["trapV2cRport"] = state ? state.trapV2cRport : undefined;
            resourceInputs["trapV2cStatus"] = state ? state.trapV2cStatus : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vdoms"] = state ? state.vdoms : undefined;
        } else {
            const args = argsOrState as CommunityArgs | undefined;
            if ((!args || args.fosid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fosid'");
            }
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["events"] = args ? args.events : undefined;
            resourceInputs["fosid"] = args ? args.fosid : undefined;
            resourceInputs["hosts"] = args ? args.hosts : undefined;
            resourceInputs["hosts6s"] = args ? args.hosts6s : undefined;
            resourceInputs["mibView"] = args ? args.mibView : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["queryV1Port"] = args ? args.queryV1Port : undefined;
            resourceInputs["queryV1Status"] = args ? args.queryV1Status : undefined;
            resourceInputs["queryV2cPort"] = args ? args.queryV2cPort : undefined;
            resourceInputs["queryV2cStatus"] = args ? args.queryV2cStatus : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["trapV1Lport"] = args ? args.trapV1Lport : undefined;
            resourceInputs["trapV1Rport"] = args ? args.trapV1Rport : undefined;
            resourceInputs["trapV1Status"] = args ? args.trapV1Status : undefined;
            resourceInputs["trapV2cLport"] = args ? args.trapV2cLport : undefined;
            resourceInputs["trapV2cRport"] = args ? args.trapV2cRport : undefined;
            resourceInputs["trapV2cStatus"] = args ? args.trapV2cStatus : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vdoms"] = args ? args.vdoms : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Community.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Community resources.
 */
export interface CommunityState {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * SNMP trap events.
     */
    events?: pulumi.Input<string>;
    /**
     * Community ID.
     */
    fosid?: pulumi.Input<number>;
    /**
     * Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
     */
    hosts?: pulumi.Input<pulumi.Input<inputs.systemsnmp.CommunityHost>[]>;
    /**
     * Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
     */
    hosts6s?: pulumi.Input<pulumi.Input<inputs.systemsnmp.CommunityHosts6>[]>;
    /**
     * SNMP access control MIB view.
     */
    mibView?: pulumi.Input<string>;
    /**
     * Community name.
     */
    name?: pulumi.Input<string>;
    /**
     * SNMP v1 query port (default = 161).
     */
    queryV1Port?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
     */
    queryV1Status?: pulumi.Input<string>;
    /**
     * SNMP v2c query port (default = 161).
     */
    queryV2cPort?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
     */
    queryV2cStatus?: pulumi.Input<string>;
    /**
     * Enable/disable this SNMP community. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * SNMP v1 trap local port (default = 162).
     */
    trapV1Lport?: pulumi.Input<number>;
    /**
     * SNMP v1 trap remote port (default = 162).
     */
    trapV1Rport?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
     */
    trapV1Status?: pulumi.Input<string>;
    /**
     * SNMP v2c trap local port (default = 162).
     */
    trapV2cLport?: pulumi.Input<number>;
    /**
     * SNMP v2c trap remote port (default = 162).
     */
    trapV2cRport?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
     */
    trapV2cStatus?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * SNMP access control VDOMs. The structure of `vdoms` block is documented below.
     */
    vdoms?: pulumi.Input<pulumi.Input<inputs.systemsnmp.CommunityVdom>[]>;
}

/**
 * The set of arguments for constructing a Community resource.
 */
export interface CommunityArgs {
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * SNMP trap events.
     */
    events?: pulumi.Input<string>;
    /**
     * Community ID.
     */
    fosid: pulumi.Input<number>;
    /**
     * Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
     */
    hosts?: pulumi.Input<pulumi.Input<inputs.systemsnmp.CommunityHost>[]>;
    /**
     * Configure IPv6 SNMP managers. The structure of `hosts6` block is documented below.
     */
    hosts6s?: pulumi.Input<pulumi.Input<inputs.systemsnmp.CommunityHosts6>[]>;
    /**
     * SNMP access control MIB view.
     */
    mibView?: pulumi.Input<string>;
    /**
     * Community name.
     */
    name?: pulumi.Input<string>;
    /**
     * SNMP v1 query port (default = 161).
     */
    queryV1Port?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v1 queries. Valid values: `enable`, `disable`.
     */
    queryV1Status?: pulumi.Input<string>;
    /**
     * SNMP v2c query port (default = 161).
     */
    queryV2cPort?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v2c queries. Valid values: `enable`, `disable`.
     */
    queryV2cStatus?: pulumi.Input<string>;
    /**
     * Enable/disable this SNMP community. Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * SNMP v1 trap local port (default = 162).
     */
    trapV1Lport?: pulumi.Input<number>;
    /**
     * SNMP v1 trap remote port (default = 162).
     */
    trapV1Rport?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v1 traps. Valid values: `enable`, `disable`.
     */
    trapV1Status?: pulumi.Input<string>;
    /**
     * SNMP v2c trap local port (default = 162).
     */
    trapV2cLport?: pulumi.Input<number>;
    /**
     * SNMP v2c trap remote port (default = 162).
     */
    trapV2cRport?: pulumi.Input<number>;
    /**
     * Enable/disable SNMP v2c traps. Valid values: `enable`, `disable`.
     */
    trapV2cStatus?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * SNMP access control VDOMs. The structure of `vdoms` block is documented below.
     */
    vdoms?: pulumi.Input<pulumi.Input<inputs.systemsnmp.CommunityVdom>[]>;
}
