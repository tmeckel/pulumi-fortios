// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const trname = new fortios.dlp.Fpdocsource("trname", {
 *     date: 1,
 *     filePath: "/",
 *     filePattern: "*",
 *     keepModified: "enable",
 *     period: "none",
 *     removeDeleted: "enable",
 *     scanOnCreation: "enable",
 *     scanSubdirectories: "enable",
 *     server: "1.1.1.1",
 *     serverType: "samba",
 *     todHour: 1,
 *     todMin: 0,
 *     username: "sgh",
 *     vdom: "mgmt",
 *     weekday: "sunday",
 * });
 * ```
 *
 * ## Import
 *
 * Dlp FpDocSource can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import fortios:dlp/fpdocsource:Fpdocsource labelname {{name}}
 * ```
 *
 *  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
 *
 * ```sh
 *  $ pulumi import fortios:dlp/fpdocsource:Fpdocsource labelname {{name}}
 * ```
 *
 *  $ unset "FORTIOS_IMPORT_TABLE"
 */
export class Fpdocsource extends pulumi.CustomResource {
    /**
     * Get an existing Fpdocsource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FpdocsourceState, opts?: pulumi.CustomResourceOptions): Fpdocsource {
        return new Fpdocsource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:dlp/fpdocsource:Fpdocsource';

    /**
     * Returns true if the given object is an instance of Fpdocsource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Fpdocsource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Fpdocsource.__pulumiType;
    }

    /**
     * Day of the month on which to scan the server (1 - 31).
     */
    public readonly date!: pulumi.Output<number>;
    /**
     * Path on the server to the fingerprint files (max 119 characters).
     */
    public readonly filePath!: pulumi.Output<string>;
    /**
     * Files matching this pattern on the server are fingerprinted. Optionally use the * and ? wildcards.
     */
    public readonly filePattern!: pulumi.Output<string>;
    /**
     * Enable so that when a file is changed on the server the FortiGate keeps the old fingerprint and adds a new fingerprint to the database. Valid values: `enable`, `disable`.
     */
    public readonly keepModified!: pulumi.Output<string>;
    /**
     * Name of the DLP fingerprint database.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password required to log into the file server.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Frequency for which the FortiGate checks the server for new or changed files. Valid values: `none`, `daily`, `weekly`, `monthly`.
     */
    public readonly period!: pulumi.Output<string>;
    /**
     * Enable to keep the fingerprint database up to date when a file is deleted from the server. Valid values: `enable`, `disable`.
     */
    public readonly removeDeleted!: pulumi.Output<string>;
    /**
     * Enable to keep the fingerprint database up to date when a file is added or changed on the server. Valid values: `enable`, `disable`.
     */
    public readonly scanOnCreation!: pulumi.Output<string>;
    /**
     * Enable/disable scanning subdirectories to find files to create fingerprints from. Valid values: `enable`, `disable`.
     */
    public readonly scanSubdirectories!: pulumi.Output<string>;
    /**
     * Select a sensitivity or threat level for matches with this fingerprint database. Add sensitivities using fp-sensitivity.
     */
    public readonly sensitivity!: pulumi.Output<string>;
    /**
     * IPv4 or IPv6 address of the server.
     */
    public readonly server!: pulumi.Output<string>;
    /**
     * Protocol used to communicate with the file server. Currently only Samba (SMB) servers are supported. Valid values: `samba`.
     */
    public readonly serverType!: pulumi.Output<string>;
    /**
     * Hour of the day on which to scan the server (0 - 23, default = 1).
     */
    public readonly todHour!: pulumi.Output<number>;
    /**
     * Minute of the hour on which to scan the server (0 - 59).
     */
    public readonly todMin!: pulumi.Output<number>;
    /**
     * User name required to log into the file server.
     */
    public readonly username!: pulumi.Output<string>;
    /**
     * Select the VDOM that can communicate with the file server. Valid values: `mgmt`, `current`.
     */
    public readonly vdom!: pulumi.Output<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    public readonly vdomparam!: pulumi.Output<string | undefined>;
    /**
     * Day of the week on which to scan the server. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    public readonly weekday!: pulumi.Output<string>;

    /**
     * Create a Fpdocsource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FpdocsourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FpdocsourceArgs | FpdocsourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FpdocsourceState | undefined;
            resourceInputs["date"] = state ? state.date : undefined;
            resourceInputs["filePath"] = state ? state.filePath : undefined;
            resourceInputs["filePattern"] = state ? state.filePattern : undefined;
            resourceInputs["keepModified"] = state ? state.keepModified : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["removeDeleted"] = state ? state.removeDeleted : undefined;
            resourceInputs["scanOnCreation"] = state ? state.scanOnCreation : undefined;
            resourceInputs["scanSubdirectories"] = state ? state.scanSubdirectories : undefined;
            resourceInputs["sensitivity"] = state ? state.sensitivity : undefined;
            resourceInputs["server"] = state ? state.server : undefined;
            resourceInputs["serverType"] = state ? state.serverType : undefined;
            resourceInputs["todHour"] = state ? state.todHour : undefined;
            resourceInputs["todMin"] = state ? state.todMin : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["vdom"] = state ? state.vdom : undefined;
            resourceInputs["vdomparam"] = state ? state.vdomparam : undefined;
            resourceInputs["weekday"] = state ? state.weekday : undefined;
        } else {
            const args = argsOrState as FpdocsourceArgs | undefined;
            if ((!args || args.server === undefined) && !opts.urn) {
                throw new Error("Missing required property 'server'");
            }
            if ((!args || args.serverType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverType'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["date"] = args ? args.date : undefined;
            resourceInputs["filePath"] = args ? args.filePath : undefined;
            resourceInputs["filePattern"] = args ? args.filePattern : undefined;
            resourceInputs["keepModified"] = args ? args.keepModified : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["removeDeleted"] = args ? args.removeDeleted : undefined;
            resourceInputs["scanOnCreation"] = args ? args.scanOnCreation : undefined;
            resourceInputs["scanSubdirectories"] = args ? args.scanSubdirectories : undefined;
            resourceInputs["sensitivity"] = args ? args.sensitivity : undefined;
            resourceInputs["server"] = args ? args.server : undefined;
            resourceInputs["serverType"] = args ? args.serverType : undefined;
            resourceInputs["todHour"] = args ? args.todHour : undefined;
            resourceInputs["todMin"] = args ? args.todMin : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["vdom"] = args ? args.vdom : undefined;
            resourceInputs["vdomparam"] = args ? args.vdomparam : undefined;
            resourceInputs["weekday"] = args ? args.weekday : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Fpdocsource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Fpdocsource resources.
 */
export interface FpdocsourceState {
    /**
     * Day of the month on which to scan the server (1 - 31).
     */
    date?: pulumi.Input<number>;
    /**
     * Path on the server to the fingerprint files (max 119 characters).
     */
    filePath?: pulumi.Input<string>;
    /**
     * Files matching this pattern on the server are fingerprinted. Optionally use the * and ? wildcards.
     */
    filePattern?: pulumi.Input<string>;
    /**
     * Enable so that when a file is changed on the server the FortiGate keeps the old fingerprint and adds a new fingerprint to the database. Valid values: `enable`, `disable`.
     */
    keepModified?: pulumi.Input<string>;
    /**
     * Name of the DLP fingerprint database.
     */
    name?: pulumi.Input<string>;
    /**
     * Password required to log into the file server.
     */
    password?: pulumi.Input<string>;
    /**
     * Frequency for which the FortiGate checks the server for new or changed files. Valid values: `none`, `daily`, `weekly`, `monthly`.
     */
    period?: pulumi.Input<string>;
    /**
     * Enable to keep the fingerprint database up to date when a file is deleted from the server. Valid values: `enable`, `disable`.
     */
    removeDeleted?: pulumi.Input<string>;
    /**
     * Enable to keep the fingerprint database up to date when a file is added or changed on the server. Valid values: `enable`, `disable`.
     */
    scanOnCreation?: pulumi.Input<string>;
    /**
     * Enable/disable scanning subdirectories to find files to create fingerprints from. Valid values: `enable`, `disable`.
     */
    scanSubdirectories?: pulumi.Input<string>;
    /**
     * Select a sensitivity or threat level for matches with this fingerprint database. Add sensitivities using fp-sensitivity.
     */
    sensitivity?: pulumi.Input<string>;
    /**
     * IPv4 or IPv6 address of the server.
     */
    server?: pulumi.Input<string>;
    /**
     * Protocol used to communicate with the file server. Currently only Samba (SMB) servers are supported. Valid values: `samba`.
     */
    serverType?: pulumi.Input<string>;
    /**
     * Hour of the day on which to scan the server (0 - 23, default = 1).
     */
    todHour?: pulumi.Input<number>;
    /**
     * Minute of the hour on which to scan the server (0 - 59).
     */
    todMin?: pulumi.Input<number>;
    /**
     * User name required to log into the file server.
     */
    username?: pulumi.Input<string>;
    /**
     * Select the VDOM that can communicate with the file server. Valid values: `mgmt`, `current`.
     */
    vdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Day of the week on which to scan the server. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    weekday?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Fpdocsource resource.
 */
export interface FpdocsourceArgs {
    /**
     * Day of the month on which to scan the server (1 - 31).
     */
    date?: pulumi.Input<number>;
    /**
     * Path on the server to the fingerprint files (max 119 characters).
     */
    filePath?: pulumi.Input<string>;
    /**
     * Files matching this pattern on the server are fingerprinted. Optionally use the * and ? wildcards.
     */
    filePattern?: pulumi.Input<string>;
    /**
     * Enable so that when a file is changed on the server the FortiGate keeps the old fingerprint and adds a new fingerprint to the database. Valid values: `enable`, `disable`.
     */
    keepModified?: pulumi.Input<string>;
    /**
     * Name of the DLP fingerprint database.
     */
    name?: pulumi.Input<string>;
    /**
     * Password required to log into the file server.
     */
    password?: pulumi.Input<string>;
    /**
     * Frequency for which the FortiGate checks the server for new or changed files. Valid values: `none`, `daily`, `weekly`, `monthly`.
     */
    period?: pulumi.Input<string>;
    /**
     * Enable to keep the fingerprint database up to date when a file is deleted from the server. Valid values: `enable`, `disable`.
     */
    removeDeleted?: pulumi.Input<string>;
    /**
     * Enable to keep the fingerprint database up to date when a file is added or changed on the server. Valid values: `enable`, `disable`.
     */
    scanOnCreation?: pulumi.Input<string>;
    /**
     * Enable/disable scanning subdirectories to find files to create fingerprints from. Valid values: `enable`, `disable`.
     */
    scanSubdirectories?: pulumi.Input<string>;
    /**
     * Select a sensitivity or threat level for matches with this fingerprint database. Add sensitivities using fp-sensitivity.
     */
    sensitivity?: pulumi.Input<string>;
    /**
     * IPv4 or IPv6 address of the server.
     */
    server: pulumi.Input<string>;
    /**
     * Protocol used to communicate with the file server. Currently only Samba (SMB) servers are supported. Valid values: `samba`.
     */
    serverType: pulumi.Input<string>;
    /**
     * Hour of the day on which to scan the server (0 - 23, default = 1).
     */
    todHour?: pulumi.Input<number>;
    /**
     * Minute of the hour on which to scan the server (0 - 59).
     */
    todMin?: pulumi.Input<number>;
    /**
     * User name required to log into the file server.
     */
    username: pulumi.Input<string>;
    /**
     * Select the VDOM that can communicate with the file server. Valid values: `mgmt`, `current`.
     */
    vdom?: pulumi.Input<string>;
    /**
     * Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
     */
    vdomparam?: pulumi.Input<string>;
    /**
     * Day of the week on which to scan the server. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
     */
    weekday?: pulumi.Input<string>;
}