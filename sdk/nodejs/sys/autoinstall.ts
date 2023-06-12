// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure USB auto installation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.sys.Autoinstall("trname", {
 *     autoInstallConfig: "enable",
 *     autoInstallImage: "enable",
 *     defaultConfigFile: "fgt_system.conf",
 *     defaultImageFile: "image.out",
 * });
 * ```
 *
 * ## Import
 *
 * System AutoInstall can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:sys/autoinstall:Autoinstall labelname SystemAutoInstall
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:sys/autoinstall:Autoinstall labelname SystemAutoInstall
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Autoinstall extends pulumi.CustomResource {
    /**
     * Get an existing Autoinstall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoinstallState, opts?: pulumi.CustomResourceOptions): Autoinstall {
        return new Autoinstall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:sys/autoinstall:Autoinstall';

    /**
     * Returns true if the given object is an instance of Autoinstall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Autoinstall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Autoinstall.__pulumiType;
    }

    /**
     * Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
     */
    public readonly autoInstallConfig!: pulumi.Output<string>;
    /**
     * Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
     */
    public readonly autoInstallImage!: pulumi.Output<string>;
    /**
     * Default config file name in USB disk.
     */
    public readonly defaultConfigFile!: pulumi.Output<string>;
    /**
     * Default image file name in USB disk.
     */
    public readonly defaultImageFile!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;

    /**
     * Create a Autoinstall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AutoinstallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoinstallArgs | AutoinstallState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoinstallState | undefined;
            resourceInputs["autoInstallConfig"] = state ? state.autoInstallConfig : undefined;
            resourceInputs["autoInstallImage"] = state ? state.autoInstallImage : undefined;
            resourceInputs["defaultConfigFile"] = state ? state.defaultConfigFile : undefined;
            resourceInputs["defaultImageFile"] = state ? state.defaultImageFile : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
        } else {
            const args = argsOrState as AutoinstallArgs | undefined;
            resourceInputs["autoInstallConfig"] = args ? args.autoInstallConfig : undefined;
            resourceInputs["autoInstallImage"] = args ? args.autoInstallImage : undefined;
            resourceInputs["defaultConfigFile"] = args ? args.defaultConfigFile : undefined;
            resourceInputs["defaultImageFile"] = args ? args.defaultImageFile : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Autoinstall.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Autoinstall resources.
 */
export interface AutoinstallState {
    /**
     * Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
     */
    autoInstallConfig?: pulumi.Input<string>;
    /**
     * Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
     */
    autoInstallImage?: pulumi.Input<string>;
    /**
     * Default config file name in USB disk.
     */
    defaultConfigFile?: pulumi.Input<string>;
    /**
     * Default image file name in USB disk.
     */
    defaultImageFile?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Autoinstall resource.
 */
export interface AutoinstallArgs {
    /**
     * Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
     */
    autoInstallConfig?: pulumi.Input<string>;
    /**
     * Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
     */
    autoInstallImage?: pulumi.Input<string>;
    /**
     * Default config file name in USB disk.
     */
    defaultConfigFile?: pulumi.Input<string>;
    /**
     * Default image file name in USB disk.
     */
    defaultImageFile?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
}
