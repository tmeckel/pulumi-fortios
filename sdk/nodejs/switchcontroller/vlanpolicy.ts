// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure VLAN policy to be applied on the managed FortiSwitch ports through port-policy. Applies to FortiOS Version `>= 6.4.0`.
 *
 * ## Import
 *
 * SwitchController VlanPolicy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:switchcontroller/vlanpolicy:Vlanpolicy labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:switchcontroller/vlanpolicy:Vlanpolicy labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Vlanpolicy extends pulumi.CustomResource {
    /**
     * Get an existing Vlanpolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VlanpolicyState, opts?: pulumi.CustomResourceOptions): Vlanpolicy {
        return new Vlanpolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/vlanpolicy:Vlanpolicy';

    /**
     * Returns true if the given object is an instance of Vlanpolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vlanpolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vlanpolicy.__pulumiType;
    }

    /**
     * Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
     */
    public readonly allowedVlans!: pulumi.Output<outputs.switchcontroller.VlanpolicyAllowedVlan[] | undefined>;
    /**
     * Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
     */
    public readonly allowedVlansAll!: pulumi.Output<string>;
    /**
     * Description for the VLAN policy.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
     */
    public readonly discardMode!: pulumi.Output<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * FortiLink interface for which this VLAN policy belongs to.
     */
    public readonly fortilink!: pulumi.Output<string>;
    /**
     * VLAN policy name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
     */
    public readonly untaggedVlans!: pulumi.Output<outputs.switchcontroller.VlanpolicyUntaggedVlan[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Native VLAN to be applied when using this VLAN policy.
     */
    public readonly vlan!: pulumi.Output<string>;

    /**
     * Create a Vlanpolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VlanpolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VlanpolicyArgs | VlanpolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VlanpolicyState | undefined;
            resourceInputs["allowedVlans"] = state ? state.allowedVlans : undefined;
            resourceInputs["allowedVlansAll"] = state ? state.allowedVlansAll : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["discardMode"] = state ? state.discardMode : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["fortilink"] = state ? state.fortilink : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["untaggedVlans"] = state ? state.untaggedVlans : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as VlanpolicyArgs | undefined;
            resourceInputs["allowedVlans"] = args ? args.allowedVlans : undefined;
            resourceInputs["allowedVlansAll"] = args ? args.allowedVlansAll : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["discardMode"] = args ? args.discardMode : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["fortilink"] = args ? args.fortilink : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["untaggedVlans"] = args ? args.untaggedVlans : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["vlan"] = args ? args.vlan : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vlanpolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vlanpolicy resources.
 */
export interface VlanpolicyState {
    /**
     * Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
     */
    allowedVlans?: pulumi.Input<pulumi.Input<inputs.switchcontroller.VlanpolicyAllowedVlan>[]>;
    /**
     * Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
     */
    allowedVlansAll?: pulumi.Input<string>;
    /**
     * Description for the VLAN policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
     */
    discardMode?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * FortiLink interface for which this VLAN policy belongs to.
     */
    fortilink?: pulumi.Input<string>;
    /**
     * VLAN policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
     */
    untaggedVlans?: pulumi.Input<pulumi.Input<inputs.switchcontroller.VlanpolicyUntaggedVlan>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Native VLAN to be applied when using this VLAN policy.
     */
    vlan?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Vlanpolicy resource.
 */
export interface VlanpolicyArgs {
    /**
     * Allowed VLANs to be applied when using this VLAN policy. The structure of `allowedVlans` block is documented below.
     */
    allowedVlans?: pulumi.Input<pulumi.Input<inputs.switchcontroller.VlanpolicyAllowedVlan>[]>;
    /**
     * Enable/disable all defined VLANs when using this VLAN policy. Valid values: `enable`, `disable`.
     */
    allowedVlansAll?: pulumi.Input<string>;
    /**
     * Description for the VLAN policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.
     */
    discardMode?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * FortiLink interface for which this VLAN policy belongs to.
     */
    fortilink?: pulumi.Input<string>;
    /**
     * VLAN policy name.
     */
    name?: pulumi.Input<string>;
    /**
     * Untagged VLANs to be applied when using this VLAN policy. The structure of `untaggedVlans` block is documented below.
     */
    untaggedVlans?: pulumi.Input<pulumi.Input<inputs.switchcontroller.VlanpolicyUntaggedVlan>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Native VLAN to be applied when using this VLAN policy.
     */
    vlan?: pulumi.Input<string>;
}
