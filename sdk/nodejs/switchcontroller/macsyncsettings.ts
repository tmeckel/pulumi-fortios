// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure global MAC synchronization settings. Applies to FortiOS Version `<= 6.2.0`.
 *
 * ## Import
 *
 * SwitchController MacSyncSettings can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:switchcontroller/macsyncsettings:Macsyncsettings labelname SwitchControllerMacSyncSettings
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:switchcontroller/macsyncsettings:Macsyncsettings labelname SwitchControllerMacSyncSettings
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Macsyncsettings extends pulumi.CustomResource {
    /**
     * Get an existing Macsyncsettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MacsyncsettingsState, opts?: pulumi.CustomResourceOptions): Macsyncsettings {
        return new Macsyncsettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:switchcontroller/macsyncsettings:Macsyncsettings';

    /**
     * Returns true if the given object is an instance of Macsyncsettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Macsyncsettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Macsyncsettings.__pulumiType;
    }

    /**
     * Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).
     */
    public readonly macSyncInterval!: pulumi.Output<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Macsyncsettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MacsyncsettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MacsyncsettingsArgs | MacsyncsettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MacsyncsettingsState | undefined;
            resourceInputs["macSyncInterval"] = state ? state.macSyncInterval : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as MacsyncsettingsArgs | undefined;
            resourceInputs["macSyncInterval"] = args ? args.macSyncInterval : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Macsyncsettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Macsyncsettings resources.
 */
export interface MacsyncsettingsState {
    /**
     * Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).
     */
    macSyncInterval?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Macsyncsettings resource.
 */
export interface MacsyncsettingsArgs {
    /**
     * Time interval between MAC synchronizations (30 - 1800 sec, default = 60, 0 = disable MAC synchronization).
     */
    macSyncInterval?: pulumi.Input<number>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
