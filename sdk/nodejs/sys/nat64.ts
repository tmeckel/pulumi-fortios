// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure NAT64. Applies to FortiOS Version `<= 7.0.0`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.sys.Nat64("trname", {
 *     alwaysSynthesizeAaaaRecord: "enable",
 *     generateIpv6FragmentHeader: "disable",
 *     nat46ForceIpv4PacketForwarding: "disable",
 *     nat64Prefix: "2001:1:2:3::/96",
 *     secondaryPrefixStatus: "disable",
 *     status: "disable",
 * });
 * ```
 *
 * ## Import
 *
 * System Nat64 can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:sys/nat64:Nat64 labelname SystemNat64
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:sys/nat64:Nat64 labelname SystemNat64
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Nat64 extends pulumi.CustomResource {
    /**
     * Get an existing Nat64 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Nat64State, opts?: pulumi.CustomResourceOptions): Nat64 {
        return new Nat64(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:sys/nat64:Nat64';

    /**
     * Returns true if the given object is an instance of Nat64.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Nat64 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Nat64.__pulumiType;
    }

    /**
     * Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
     */
    public readonly alwaysSynthesizeAaaaRecord!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
     */
    public readonly generateIpv6FragmentHeader!: pulumi.Output<string>;
    /**
     * Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
     */
    public readonly nat46ForceIpv4PacketForwarding!: pulumi.Output<string>;
    /**
     * NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
     */
    public readonly nat64Prefix!: pulumi.Output<string>;
    /**
     * Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
     */
    public readonly secondaryPrefixStatus!: pulumi.Output<string>;
    /**
     * Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
     */
    public readonly secondaryPrefixes!: pulumi.Output<outputs.sys.Nat64SecondaryPrefix[] | undefined>;
    /**
     * Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
     */
    public readonly status!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Nat64 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Nat64Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Nat64Args | Nat64State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Nat64State | undefined;
            resourceInputs["alwaysSynthesizeAaaaRecord"] = state ? state.alwaysSynthesizeAaaaRecord : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["generateIpv6FragmentHeader"] = state ? state.generateIpv6FragmentHeader : undefined;
            resourceInputs["nat46ForceIpv4PacketForwarding"] = state ? state.nat46ForceIpv4PacketForwarding : undefined;
            resourceInputs["nat64Prefix"] = state ? state.nat64Prefix : undefined;
            resourceInputs["secondaryPrefixStatus"] = state ? state.secondaryPrefixStatus : undefined;
            resourceInputs["secondaryPrefixes"] = state ? state.secondaryPrefixes : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as Nat64Args | undefined;
            if ((!args || args.nat64Prefix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nat64Prefix'");
            }
            resourceInputs["alwaysSynthesizeAaaaRecord"] = args ? args.alwaysSynthesizeAaaaRecord : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["generateIpv6FragmentHeader"] = args ? args.generateIpv6FragmentHeader : undefined;
            resourceInputs["nat46ForceIpv4PacketForwarding"] = args ? args.nat46ForceIpv4PacketForwarding : undefined;
            resourceInputs["nat64Prefix"] = args ? args.nat64Prefix : undefined;
            resourceInputs["secondaryPrefixStatus"] = args ? args.secondaryPrefixStatus : undefined;
            resourceInputs["secondaryPrefixes"] = args ? args.secondaryPrefixes : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Nat64.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Nat64 resources.
 */
export interface Nat64State {
    /**
     * Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
     */
    alwaysSynthesizeAaaaRecord?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
     */
    generateIpv6FragmentHeader?: pulumi.Input<string>;
    /**
     * Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
     */
    nat46ForceIpv4PacketForwarding?: pulumi.Input<string>;
    /**
     * NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
     */
    nat64Prefix?: pulumi.Input<string>;
    /**
     * Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
     */
    secondaryPrefixStatus?: pulumi.Input<string>;
    /**
     * Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
     */
    secondaryPrefixes?: pulumi.Input<pulumi.Input<inputs.sys.Nat64SecondaryPrefix>[]>;
    /**
     * Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Nat64 resource.
 */
export interface Nat64Args {
    /**
     * Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
     */
    alwaysSynthesizeAaaaRecord?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Enable/disable IPv6 fragment header generation. Valid values: `enable`, `disable`.
     */
    generateIpv6FragmentHeader?: pulumi.Input<string>;
    /**
     * Enable/disable mandatory IPv4 packet forwarding in nat46. Valid values: `enable`, `disable`.
     */
    nat46ForceIpv4PacketForwarding?: pulumi.Input<string>;
    /**
     * NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
     */
    nat64Prefix: pulumi.Input<string>;
    /**
     * Enable/disable secondary NAT64 prefix. Valid values: `enable`, `disable`.
     */
    secondaryPrefixStatus?: pulumi.Input<string>;
    /**
     * Secondary NAT64 prefix. The structure of `secondaryPrefix` block is documented below.
     */
    secondaryPrefixes?: pulumi.Input<pulumi.Input<inputs.sys.Nat64SecondaryPrefix>[]>;
    /**
     * Enable/disable NAT64 (default = disable). Valid values: `enable`, `disable`.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
