// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure FortiSwitch LLDP profiles.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.switchcontroller.Lldpprofile("trname", {
 *     autoIsl: "enable",
 *     autoIslHelloTimer: 3,
 *     autoIslPortGroup: 0,
 *     autoIslReceiveTimeout: 60,
 *     medTlvs: "inventory-management network-policy",
 * });
 * ```
 *
 * ## Import
 *
 * SwitchController LldpProfile can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:switchcontroller/lldpprofile:Lldpprofile labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:switchcontroller/lldpprofile:Lldpprofile labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Lldpprofile extends pulumi.CustomResource {
    /**
     * Get an existing Lldpprofile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LldpprofileState, opts?: pulumi.CustomResourceOptions): Lldpprofile {
        return new Lldpprofile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/lldpprofile:Lldpprofile';

    /**
     * Returns true if the given object is an instance of Lldpprofile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Lldpprofile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Lldpprofile.__pulumiType;
    }

    /**
     * Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.
     */
    public readonly autoIsl!: pulumi.Output<string>;
    /**
     * Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
     */
    public readonly autoIslHelloTimer!: pulumi.Output<number>;
    /**
     * Auto inter-switch LAG port group ID (0 - 9).
     */
    public readonly autoIslPortGroup!: pulumi.Output<number>;
    /**
     * Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
     */
    public readonly autoIslReceiveTimeout!: pulumi.Output<number>;
    /**
     * Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.
     */
    public readonly autoMclagIcl!: pulumi.Output<string>;
    /**
     * Configuration method to edit custom TLV entries. The structure of `customTlvs` block is documented below.
     */
    public readonly customTlvs!: pulumi.Output<outputs.switchcontroller.LldpprofileCustomTlv[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `medLocationService` block is documented below.
     */
    public readonly medLocationServices!: pulumi.Output<outputs.switchcontroller.LldpprofileMedLocationService[] | undefined>;
    /**
     * Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `medNetworkPolicy` block is documented below.
     */
    public readonly medNetworkPolicies!: pulumi.Output<outputs.switchcontroller.LldpprofileMedNetworkPolicy[] | undefined>;
    /**
     * Transmitted LLDP-MED TLVs (type-length-value descriptions): inventory management TLV and/or network policy TLV.
     */
    public readonly medTlvs!: pulumi.Output<string>;
    /**
     * Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.
     */
    public readonly n8021Tlvs!: pulumi.Output<string>;
    /**
     * Transmitted IEEE 802.3 TLVs.
     */
    public readonly n8023Tlvs!: pulumi.Output<string>;
    /**
     * Profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Lldpprofile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LldpprofileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LldpprofileArgs | LldpprofileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LldpprofileState | undefined;
            resourceInputs["autoIsl"] = state ? state.autoIsl : undefined;
            resourceInputs["autoIslHelloTimer"] = state ? state.autoIslHelloTimer : undefined;
            resourceInputs["autoIslPortGroup"] = state ? state.autoIslPortGroup : undefined;
            resourceInputs["autoIslReceiveTimeout"] = state ? state.autoIslReceiveTimeout : undefined;
            resourceInputs["autoMclagIcl"] = state ? state.autoMclagIcl : undefined;
            resourceInputs["customTlvs"] = state ? state.customTlvs : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["medLocationServices"] = state ? state.medLocationServices : undefined;
            resourceInputs["medNetworkPolicies"] = state ? state.medNetworkPolicies : undefined;
            resourceInputs["medTlvs"] = state ? state.medTlvs : undefined;
            resourceInputs["n8021Tlvs"] = state ? state.n8021Tlvs : undefined;
            resourceInputs["n8023Tlvs"] = state ? state.n8023Tlvs : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as LldpprofileArgs | undefined;
            resourceInputs["autoIsl"] = args ? args.autoIsl : undefined;
            resourceInputs["autoIslHelloTimer"] = args ? args.autoIslHelloTimer : undefined;
            resourceInputs["autoIslPortGroup"] = args ? args.autoIslPortGroup : undefined;
            resourceInputs["autoIslReceiveTimeout"] = args ? args.autoIslReceiveTimeout : undefined;
            resourceInputs["autoMclagIcl"] = args ? args.autoMclagIcl : undefined;
            resourceInputs["customTlvs"] = args ? args.customTlvs : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["medLocationServices"] = args ? args.medLocationServices : undefined;
            resourceInputs["medNetworkPolicies"] = args ? args.medNetworkPolicies : undefined;
            resourceInputs["medTlvs"] = args ? args.medTlvs : undefined;
            resourceInputs["n8021Tlvs"] = args ? args.n8021Tlvs : undefined;
            resourceInputs["n8023Tlvs"] = args ? args.n8023Tlvs : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Lldpprofile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Lldpprofile resources.
 */
export interface LldpprofileState {
    /**
     * Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.
     */
    autoIsl?: pulumi.Input<string>;
    /**
     * Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
     */
    autoIslHelloTimer?: pulumi.Input<number>;
    /**
     * Auto inter-switch LAG port group ID (0 - 9).
     */
    autoIslPortGroup?: pulumi.Input<number>;
    /**
     * Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
     */
    autoIslReceiveTimeout?: pulumi.Input<number>;
    /**
     * Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.
     */
    autoMclagIcl?: pulumi.Input<string>;
    /**
     * Configuration method to edit custom TLV entries. The structure of `customTlvs` block is documented below.
     */
    customTlvs?: pulumi.Input<pulumi.Input<inputs.switchcontroller.LldpprofileCustomTlv>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `medLocationService` block is documented below.
     */
    medLocationServices?: pulumi.Input<pulumi.Input<inputs.switchcontroller.LldpprofileMedLocationService>[]>;
    /**
     * Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `medNetworkPolicy` block is documented below.
     */
    medNetworkPolicies?: pulumi.Input<pulumi.Input<inputs.switchcontroller.LldpprofileMedNetworkPolicy>[]>;
    /**
     * Transmitted LLDP-MED TLVs (type-length-value descriptions): inventory management TLV and/or network policy TLV.
     */
    medTlvs?: pulumi.Input<string>;
    /**
     * Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.
     */
    n8021Tlvs?: pulumi.Input<string>;
    /**
     * Transmitted IEEE 802.3 TLVs.
     */
    n8023Tlvs?: pulumi.Input<string>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Lldpprofile resource.
 */
export interface LldpprofileArgs {
    /**
     * Enable/disable auto inter-switch LAG. Valid values: `disable`, `enable`.
     */
    autoIsl?: pulumi.Input<string>;
    /**
     * Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).
     */
    autoIslHelloTimer?: pulumi.Input<number>;
    /**
     * Auto inter-switch LAG port group ID (0 - 9).
     */
    autoIslPortGroup?: pulumi.Input<number>;
    /**
     * Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).
     */
    autoIslReceiveTimeout?: pulumi.Input<number>;
    /**
     * Enable/disable MCLAG inter chassis link. Valid values: `disable`, `enable`.
     */
    autoMclagIcl?: pulumi.Input<string>;
    /**
     * Configuration method to edit custom TLV entries. The structure of `customTlvs` block is documented below.
     */
    customTlvs?: pulumi.Input<pulumi.Input<inputs.switchcontroller.LldpprofileCustomTlv>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories. The structure of `medLocationService` block is documented below.
     */
    medLocationServices?: pulumi.Input<pulumi.Input<inputs.switchcontroller.LldpprofileMedLocationService>[]>;
    /**
     * Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories. The structure of `medNetworkPolicy` block is documented below.
     */
    medNetworkPolicies?: pulumi.Input<pulumi.Input<inputs.switchcontroller.LldpprofileMedNetworkPolicy>[]>;
    /**
     * Transmitted LLDP-MED TLVs (type-length-value descriptions): inventory management TLV and/or network policy TLV.
     */
    medTlvs?: pulumi.Input<string>;
    /**
     * Transmitted IEEE 802.1 TLVs. Valid values: `port-vlan-id`.
     */
    n8021Tlvs?: pulumi.Input<string>;
    /**
     * Transmitted IEEE 802.3 TLVs.
     */
    n8023Tlvs?: pulumi.Input<string>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
