// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure system-wide switch controller settings.
 *
 * ## Import
 *
 * SwitchController System can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:switchcontroller/system:System labelname SwitchControllerSystem
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:switchcontroller/system:System labelname SwitchControllerSystem
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class System extends pulumi.CustomResource {
    /**
     * Get an existing System resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SystemState, opts?: pulumi.CustomResourceOptions): System {
        return new System(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/system:System';

    /**
     * Returns true if the given object is an instance of System.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is System {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === System.__pulumiType;
    }

    /**
     * Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
     */
    public readonly dataSyncInterval!: pulumi.Output<number>;
    /**
     * Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
     */
    public readonly dynamicPeriodicInterval!: pulumi.Output<number>;
    /**
     * MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
     */
    public readonly iotHoldoff!: pulumi.Output<number>;
    /**
     * MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
     */
    public readonly iotMacIdle!: pulumi.Output<number>;
    /**
     * IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
     */
    public readonly iotScanInterval!: pulumi.Output<number>;
    /**
     * MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
     */
    public readonly iotWeightThreshold!: pulumi.Output<number>;
    /**
     * Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
     */
    public readonly nacPeriodicInterval!: pulumi.Output<number>;
    /**
     * Maximum number of parallel processes (1 - 300, default = 1).
     */
    public readonly parallelProcess!: pulumi.Output<number>;
    /**
     * Enable/disable parallel process override. Valid values: `disable`, `enable`.
     */
    public readonly parallelProcessOverride!: pulumi.Output<string>;
    /**
     * Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
     */
    public readonly tunnelMode!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a System resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SystemArgs | SystemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SystemState | undefined;
            resourceInputs["dataSyncInterval"] = state ? state.dataSyncInterval : undefined;
            resourceInputs["dynamicPeriodicInterval"] = state ? state.dynamicPeriodicInterval : undefined;
            resourceInputs["iotHoldoff"] = state ? state.iotHoldoff : undefined;
            resourceInputs["iotMacIdle"] = state ? state.iotMacIdle : undefined;
            resourceInputs["iotScanInterval"] = state ? state.iotScanInterval : undefined;
            resourceInputs["iotWeightThreshold"] = state ? state.iotWeightThreshold : undefined;
            resourceInputs["nacPeriodicInterval"] = state ? state.nacPeriodicInterval : undefined;
            resourceInputs["parallelProcess"] = state ? state.parallelProcess : undefined;
            resourceInputs["parallelProcessOverride"] = state ? state.parallelProcessOverride : undefined;
            resourceInputs["tunnelMode"] = state ? state.tunnelMode : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as SystemArgs | undefined;
            resourceInputs["dataSyncInterval"] = args ? args.dataSyncInterval : undefined;
            resourceInputs["dynamicPeriodicInterval"] = args ? args.dynamicPeriodicInterval : undefined;
            resourceInputs["iotHoldoff"] = args ? args.iotHoldoff : undefined;
            resourceInputs["iotMacIdle"] = args ? args.iotMacIdle : undefined;
            resourceInputs["iotScanInterval"] = args ? args.iotScanInterval : undefined;
            resourceInputs["iotWeightThreshold"] = args ? args.iotWeightThreshold : undefined;
            resourceInputs["nacPeriodicInterval"] = args ? args.nacPeriodicInterval : undefined;
            resourceInputs["parallelProcess"] = args ? args.parallelProcess : undefined;
            resourceInputs["parallelProcessOverride"] = args ? args.parallelProcessOverride : undefined;
            resourceInputs["tunnelMode"] = args ? args.tunnelMode : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(System.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering System resources.
 */
export interface SystemState {
    /**
     * Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
     */
    dataSyncInterval?: pulumi.Input<number>;
    /**
     * Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
     */
    dynamicPeriodicInterval?: pulumi.Input<number>;
    /**
     * MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
     */
    iotHoldoff?: pulumi.Input<number>;
    /**
     * MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
     */
    iotMacIdle?: pulumi.Input<number>;
    /**
     * IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
     */
    iotScanInterval?: pulumi.Input<number>;
    /**
     * MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
     */
    iotWeightThreshold?: pulumi.Input<number>;
    /**
     * Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
     */
    nacPeriodicInterval?: pulumi.Input<number>;
    /**
     * Maximum number of parallel processes (1 - 300, default = 1).
     */
    parallelProcess?: pulumi.Input<number>;
    /**
     * Enable/disable parallel process override. Valid values: `disable`, `enable`.
     */
    parallelProcessOverride?: pulumi.Input<string>;
    /**
     * Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
     */
    tunnelMode?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a System resource.
 */
export interface SystemArgs {
    /**
     * Time interval between collection of switch data (30 - 1800 sec, default = 60, 0 = disable).
     */
    dataSyncInterval?: pulumi.Input<number>;
    /**
     * Periodic time interval to run Dynamic port policy engine (5 - 60 sec, default = 15).
     */
    dynamicPeriodicInterval?: pulumi.Input<number>;
    /**
     * MAC entry's creation time. Time must be greater than this value for an entry to be created (default = 5 mins).
     */
    iotHoldoff?: pulumi.Input<number>;
    /**
     * MAC entry's idle time. MAC entry is removed after this value (default = 1440 mins).
     */
    iotMacIdle?: pulumi.Input<number>;
    /**
     * IoT scan interval (2 - 4294967295 mins, default = 60 mins, 0 = disable).
     */
    iotScanInterval?: pulumi.Input<number>;
    /**
     * MAC entry's confidence value. Value is re-queried when below this value (default = 1, 0 = disable).
     */
    iotWeightThreshold?: pulumi.Input<number>;
    /**
     * Periodic time interval to run NAC engine (5 - 60 sec, default = 15).
     */
    nacPeriodicInterval?: pulumi.Input<number>;
    /**
     * Maximum number of parallel processes (1 - 300, default = 1).
     */
    parallelProcess?: pulumi.Input<number>;
    /**
     * Enable/disable parallel process override. Valid values: `disable`, `enable`.
     */
    parallelProcessOverride?: pulumi.Input<string>;
    /**
     * Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.
     */
    tunnelMode?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}