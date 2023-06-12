// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configure FortiSwitch QoS egress queue policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.switchcontrollerqos.Queuepolicy("trname", {
 *     rateBy: "kbps",
 *     schedule: "round-robin",
 * });
 * ```
 *
 * ## Import
 *
 * SwitchControllerQos QueuePolicy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:switchcontrollerqos/queuepolicy:Queuepolicy labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:switchcontrollerqos/queuepolicy:Queuepolicy labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Queuepolicy extends pulumi.CustomResource {
    /**
     * Get an existing Queuepolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueuepolicyState, opts?: pulumi.CustomResourceOptions): Queuepolicy {
        return new Queuepolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontrollerqos/queuepolicy:Queuepolicy';

    /**
     * Returns true if the given object is an instance of Queuepolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Queuepolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Queuepolicy.__pulumiType;
    }

    /**
     * COS queue configuration. The structure of `cosQueue` block is documented below.
     */
    public readonly cosQueues!: pulumi.Output<outputs.switchcontrollerqos.QueuepolicyCosQueue[] | undefined>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    public readonly dynamicSortSubtable!: pulumi.Output<string | undefined>;
    /**
     * QoS policy name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
     */
    public readonly rateBy!: pulumi.Output<string>;
    /**
     * COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Queuepolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueuepolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueuepolicyArgs | QueuepolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QueuepolicyState | undefined;
            resourceInputs["cosQueues"] = state ? state.cosQueues : undefined;
            resourceInputs["dynamicSortSubtable"] = state ? state.dynamicSortSubtable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rateBy"] = state ? state.rateBy : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as QueuepolicyArgs | undefined;
            if ((!args || args.rateBy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rateBy'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            resourceInputs["cosQueues"] = args ? args.cosQueues : undefined;
            resourceInputs["dynamicSortSubtable"] = args ? args.dynamicSortSubtable : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rateBy"] = args ? args.rateBy : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Queuepolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Queuepolicy resources.
 */
export interface QueuepolicyState {
    /**
     * COS queue configuration. The structure of `cosQueue` block is documented below.
     */
    cosQueues?: pulumi.Input<pulumi.Input<inputs.switchcontrollerqos.QueuepolicyCosQueue>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * QoS policy name
     */
    name?: pulumi.Input<string>;
    /**
     * COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
     */
    rateBy?: pulumi.Input<string>;
    /**
     * COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
     */
    schedule?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Queuepolicy resource.
 */
export interface QueuepolicyArgs {
    /**
     * COS queue configuration. The structure of `cosQueue` block is documented below.
     */
    cosQueues?: pulumi.Input<pulumi.Input<inputs.switchcontrollerqos.QueuepolicyCosQueue>[]>;
    /**
     * Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
     */
    dynamicSortSubtable?: pulumi.Input<string>;
    /**
     * QoS policy name
     */
    name?: pulumi.Input<string>;
    /**
     * COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.
     */
    rateBy: pulumi.Input<string>;
    /**
     * COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.
     */
    schedule: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}