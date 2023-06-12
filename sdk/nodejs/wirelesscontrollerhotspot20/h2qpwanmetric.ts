// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure WAN metrics.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.wirelesscontrollerhotspot20.H2qpwanmetric("trname", {
 *     downlinkLoad: 0,
 *     downlinkSpeed: 2400,
 *     linkAtCapacity: "disable",
 *     linkStatus: "up",
 *     loadMeasurementDuration: 0,
 *     symmetricWanLink: "symmetric",
 *     uplinkLoad: 0,
 *     uplinkSpeed: 2400,
 * });
 * ```
 *
 * ## Import
 *
 * WirelessControllerHotspot20 H2QpWanMetric can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:wirelesscontrollerhotspot20/h2qpwanmetric:H2qpwanmetric labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:wirelesscontrollerhotspot20/h2qpwanmetric:H2qpwanmetric labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class H2qpwanmetric extends pulumi.CustomResource {
    /**
     * Get an existing H2qpwanmetric resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: H2qpwanmetricState, opts?: pulumi.CustomResourceOptions): H2qpwanmetric {
        return new H2qpwanmetric(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:wirelesscontrollerhotspot20/h2qpwanmetric:H2qpwanmetric';

    /**
     * Returns true if the given object is an instance of H2qpwanmetric.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is H2qpwanmetric {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === H2qpwanmetric.__pulumiType;
    }

    /**
     * Downlink load.
     */
    public readonly downlinkLoad!: pulumi.Output<number>;
    /**
     * Downlink speed (in kilobits/s).
     */
    public readonly downlinkSpeed!: pulumi.Output<number>;
    /**
     * Link at capacity. Valid values: `enable`, `disable`.
     */
    public readonly linkAtCapacity!: pulumi.Output<string>;
    /**
     * Link status. Valid values: `up`, `down`, `in-test`.
     */
    public readonly linkStatus!: pulumi.Output<string>;
    /**
     * Load measurement duration (in tenths of a second).
     */
    public readonly loadMeasurementDuration!: pulumi.Output<number>;
    /**
     * WAN metric name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
     */
    public readonly symmetricWanLink!: pulumi.Output<string>;
    /**
     * Uplink load.
     */
    public readonly uplinkLoad!: pulumi.Output<number>;
    /**
     * Uplink speed (in kilobits/s).
     */
    public readonly uplinkSpeed!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a H2qpwanmetric resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: H2qpwanmetricArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: H2qpwanmetricArgs | H2qpwanmetricState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as H2qpwanmetricState | undefined;
            resourceInputs["downlinkLoad"] = state ? state.downlinkLoad : undefined;
            resourceInputs["downlinkSpeed"] = state ? state.downlinkSpeed : undefined;
            resourceInputs["linkAtCapacity"] = state ? state.linkAtCapacity : undefined;
            resourceInputs["linkStatus"] = state ? state.linkStatus : undefined;
            resourceInputs["loadMeasurementDuration"] = state ? state.loadMeasurementDuration : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["symmetricWanLink"] = state ? state.symmetricWanLink : undefined;
            resourceInputs["uplinkLoad"] = state ? state.uplinkLoad : undefined;
            resourceInputs["uplinkSpeed"] = state ? state.uplinkSpeed : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as H2qpwanmetricArgs | undefined;
            resourceInputs["downlinkLoad"] = args ? args.downlinkLoad : undefined;
            resourceInputs["downlinkSpeed"] = args ? args.downlinkSpeed : undefined;
            resourceInputs["linkAtCapacity"] = args ? args.linkAtCapacity : undefined;
            resourceInputs["linkStatus"] = args ? args.linkStatus : undefined;
            resourceInputs["loadMeasurementDuration"] = args ? args.loadMeasurementDuration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["symmetricWanLink"] = args ? args.symmetricWanLink : undefined;
            resourceInputs["uplinkLoad"] = args ? args.uplinkLoad : undefined;
            resourceInputs["uplinkSpeed"] = args ? args.uplinkSpeed : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(H2qpwanmetric.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering H2qpwanmetric resources.
 */
export interface H2qpwanmetricState {
    /**
     * Downlink load.
     */
    downlinkLoad?: pulumi.Input<number>;
    /**
     * Downlink speed (in kilobits/s).
     */
    downlinkSpeed?: pulumi.Input<number>;
    /**
     * Link at capacity. Valid values: `enable`, `disable`.
     */
    linkAtCapacity?: pulumi.Input<string>;
    /**
     * Link status. Valid values: `up`, `down`, `in-test`.
     */
    linkStatus?: pulumi.Input<string>;
    /**
     * Load measurement duration (in tenths of a second).
     */
    loadMeasurementDuration?: pulumi.Input<number>;
    /**
     * WAN metric name.
     */
    name?: pulumi.Input<string>;
    /**
     * WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
     */
    symmetricWanLink?: pulumi.Input<string>;
    /**
     * Uplink load.
     */
    uplinkLoad?: pulumi.Input<number>;
    /**
     * Uplink speed (in kilobits/s).
     */
    uplinkSpeed?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a H2qpwanmetric resource.
 */
export interface H2qpwanmetricArgs {
    /**
     * Downlink load.
     */
    downlinkLoad?: pulumi.Input<number>;
    /**
     * Downlink speed (in kilobits/s).
     */
    downlinkSpeed?: pulumi.Input<number>;
    /**
     * Link at capacity. Valid values: `enable`, `disable`.
     */
    linkAtCapacity?: pulumi.Input<string>;
    /**
     * Link status. Valid values: `up`, `down`, `in-test`.
     */
    linkStatus?: pulumi.Input<string>;
    /**
     * Load measurement duration (in tenths of a second).
     */
    loadMeasurementDuration?: pulumi.Input<number>;
    /**
     * WAN metric name.
     */
    name?: pulumi.Input<string>;
    /**
     * WAN link symmetry. Valid values: `symmetric`, `asymmetric`.
     */
    symmetricWanLink?: pulumi.Input<string>;
    /**
     * Uplink load.
     */
    uplinkLoad?: pulumi.Input<number>;
    /**
     * Uplink speed (in kilobits/s).
     */
    uplinkSpeed?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}