// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure device groups. Applies to FortiOS Version `<= 6.2.0`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trnames12 = new fortios.user.Device("trnames12", {
 *     alias: "user_devices2",
 *     category: "amazon-device",
 *     mac: "08:00:20:0a:1c:1d",
 *     type: "unknown",
 * });
 * const trname = new fortios.user.Devicegroup("trname", {members: [{
 *     name: trnames12.alias,
 * }]});
 * ```
 *
 * ## Import
 *
 * User DeviceGroup can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:user/devicegroup:Devicegroup labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:user/devicegroup:Devicegroup labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Devicegroup extends pulumi.CustomResource {
    /**
     * Get an existing Devicegroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DevicegroupState, opts?: pulumi.CustomResourceOptions): Devicegroup {
        return new Devicegroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:user/devicegroup:Devicegroup';

    /**
     * Returns true if the given object is an instance of Devicegroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Devicegroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Devicegroup.__pulumiType;
    }

    /**
     * Comment.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * Device group member. The structure of `member` block is documented below.
     */
    public readonly members!: pulumi.Output<outputs.user.DevicegroupMember[] | undefined>;
    /**
     * Device group name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    public readonly taggings!: pulumi.Output<outputs.user.DevicegroupTagging[] | undefined>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Devicegroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DevicegroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DevicegroupArgs | DevicegroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DevicegroupState | undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["taggings"] = state ? state.taggings : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as DevicegroupArgs | undefined;
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["taggings"] = args ? args.taggings : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Devicegroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Devicegroup resources.
 */
export interface DevicegroupState {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Device group member. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.user.DevicegroupMember>[]>;
    /**
     * Device group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.user.DevicegroupTagging>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Devicegroup resource.
 */
export interface DevicegroupArgs {
    /**
     * Comment.
     */
    comment?: pulumi.Input<string>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * Device group member. The structure of `member` block is documented below.
     */
    members?: pulumi.Input<pulumi.Input<inputs.user.DevicegroupMember>[]>;
    /**
     * Device group name.
     */
    name?: pulumi.Input<string>;
    /**
     * Config object tagging. The structure of `tagging` block is documented below.
     */
    taggings?: pulumi.Input<pulumi.Input<inputs.user.DevicegroupTagging>[]>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
