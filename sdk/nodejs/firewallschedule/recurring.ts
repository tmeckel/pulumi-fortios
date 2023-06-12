// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Recurring schedule configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.firewallschedule.Recurring("trname", {
 *     color: 0,
 *     day: "sunday",
 *     end: "00:00",
 *     start: "00:00",
 * });
 * ```
 *
 * ## Import
 *
 * FirewallSchedule Recurring can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:firewallschedule/recurring:Recurring labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:firewallschedule/recurring:Recurring labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Recurring extends pulumi.CustomResource {
    /**
     * Get an existing Recurring resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecurringState, opts?: pulumi.CustomResourceOptions): Recurring {
        return new Recurring(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:firewallschedule/recurring:Recurring';

    /**
     * Returns true if the given object is an instance of Recurring.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Recurring {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Recurring.__pulumiType;
    }

    /**
     * Color of icon on the GUI.
     */
    public readonly color!: pulumi.Output<number>;
    /**
     * One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
     */
    public readonly day!: pulumi.Output<string>;
    /**
     * Time of day to end the schedule, format hh:mm.
     */
    public readonly end!: pulumi.Output<string>;
    /**
     * Security Fabric global object setting. Valid values: `enable`, `disable`.
     */
    public readonly fabricObject!: pulumi.Output<string>;
    /**
     * Recurring schedule name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Time of day to start the schedule, format hh:mm.
     */
    public readonly start!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Recurring resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecurringArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecurringArgs | RecurringState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RecurringState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["day"] = state ? state.day : undefined;
            resourceInputs["end"] = state ? state.end : undefined;
            resourceInputs["fabricObject"] = state ? state.fabricObject : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["start"] = state ? state.start : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as RecurringArgs | undefined;
            if ((!args || args.end === undefined) && !opts.urn) {
                throw new Error("Missing required property 'end'");
            }
            if ((!args || args.start === undefined) && !opts.urn) {
                throw new Error("Missing required property 'start'");
            }
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["day"] = args ? args.day : undefined;
            resourceInputs["end"] = args ? args.end : undefined;
            resourceInputs["fabricObject"] = args ? args.fabricObject : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["start"] = args ? args.start : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Recurring.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Recurring resources.
 */
export interface RecurringState {
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
     */
    day?: pulumi.Input<string>;
    /**
     * Time of day to end the schedule, format hh:mm.
     */
    end?: pulumi.Input<string>;
    /**
     * Security Fabric global object setting. Valid values: `enable`, `disable`.
     */
    fabricObject?: pulumi.Input<string>;
    /**
     * Recurring schedule name.
     */
    name?: pulumi.Input<string>;
    /**
     * Time of day to start the schedule, format hh:mm.
     */
    start?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Recurring resource.
 */
export interface RecurringArgs {
    /**
     * Color of icon on the GUI.
     */
    color?: pulumi.Input<number>;
    /**
     * One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.
     */
    day?: pulumi.Input<string>;
    /**
     * Time of day to end the schedule, format hh:mm.
     */
    end: pulumi.Input<string>;
    /**
     * Security Fabric global object setting. Valid values: `enable`, `disable`.
     */
    fabricObject?: pulumi.Input<string>;
    /**
     * Recurring schedule name.
     */
    name?: pulumi.Input<string>;
    /**
     * Time of day to start the schedule, format hh:mm.
     */
    start: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
