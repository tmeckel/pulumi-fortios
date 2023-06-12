// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to configure access profiles of FortiOS.
 *
 * !> **Warning:** The resource will be deprecated and replaced by new resource `fortios.sys.Accprofile`, we recommend that you use the new resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fortios from "@pulumiverse/fortios";
 *
 * const test1 = new fortios.sys.AdminProfiles("test1", {
 *     admintimeoutOverride: "disable",
 *     authgrp: "none",
 *     comments: "test",
 *     ftviewgrp: "read",
 *     fwgrp: "none",
 *     loggrp: "none",
 *     netgrp: "none",
 *     scope: "vdom",
 *     secfabgrp: "read-write",
 *     sysgrp: "read",
 *     utmgrp: "none",
 *     vpngrp: "none",
 *     wanoptgrp: "none",
 *     wifi: "none",
 * });
 * ```
 */
export class AdminProfiles extends pulumi.CustomResource {
    /**
     * Get an existing AdminProfiles resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdminProfilesState, opts?: pulumi.CustomResourceOptions): AdminProfiles {
        return new AdminProfiles(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'fortios:sys/adminProfiles:AdminProfiles';

    /**
     * Returns true if the given object is an instance of AdminProfiles.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdminProfiles {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdminProfiles.__pulumiType;
    }

    /**
     * Enable/disable overriding the global administrator idle timeout.
     */
    public readonly admintimeoutOverride!: pulumi.Output<string>;
    /**
     * Administrator access to Users and Devices.
     */
    public readonly authgrp!: pulumi.Output<string>;
    /**
     * Comment.
     */
    public readonly comments!: pulumi.Output<string | undefined>;
    /**
     * FortiView.
     */
    public readonly ftviewgrp!: pulumi.Output<string>;
    /**
     * Administrator access to the Firewall configuration.
     */
    public readonly fwgrp!: pulumi.Output<string>;
    /**
     * Administrator access to Logging and Reporting including viewing log messages.
     */
    public readonly loggrp!: pulumi.Output<string>;
    /**
     * Profile name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Network Configuration.
     */
    public readonly netgrp!: pulumi.Output<string>;
    /**
     * Scope of admin access.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * Security Fabric.
     */
    public readonly secfabgrp!: pulumi.Output<string>;
    /**
     * System Configuration.
     */
    public readonly sysgrp!: pulumi.Output<string>;
    /**
     * Administrator access to Security Profiles.
     */
    public readonly utmgrp!: pulumi.Output<string>;
    /**
     * Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
     */
    public readonly vpngrp!: pulumi.Output<string>;
    /**
     * Administrator access to WAN Opt & Cache.
     */
    public readonly wanoptgrp!: pulumi.Output<string>;
    /**
     * Administrator access to the WiFi controller and Switch controller.
     */
    public readonly wifi!: pulumi.Output<string>;

    /**
     * Create a AdminProfiles resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AdminProfilesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdminProfilesArgs | AdminProfilesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdminProfilesState | undefined;
            resourceInputs["admintimeoutOverride"] = state ? state.admintimeoutOverride : undefined;
            resourceInputs["authgrp"] = state ? state.authgrp : undefined;
            resourceInputs["comments"] = state ? state.comments : undefined;
            resourceInputs["ftviewgrp"] = state ? state.ftviewgrp : undefined;
            resourceInputs["fwgrp"] = state ? state.fwgrp : undefined;
            resourceInputs["loggrp"] = state ? state.loggrp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["netgrp"] = state ? state.netgrp : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["secfabgrp"] = state ? state.secfabgrp : undefined;
            resourceInputs["sysgrp"] = state ? state.sysgrp : undefined;
            resourceInputs["utmgrp"] = state ? state.utmgrp : undefined;
            resourceInputs["vpngrp"] = state ? state.vpngrp : undefined;
            resourceInputs["wanoptgrp"] = state ? state.wanoptgrp : undefined;
            resourceInputs["wifi"] = state ? state.wifi : undefined;
        } else {
            const args = argsOrState as AdminProfilesArgs | undefined;
            resourceInputs["admintimeoutOverride"] = args ? args.admintimeoutOverride : undefined;
            resourceInputs["authgrp"] = args ? args.authgrp : undefined;
            resourceInputs["comments"] = args ? args.comments : undefined;
            resourceInputs["ftviewgrp"] = args ? args.ftviewgrp : undefined;
            resourceInputs["fwgrp"] = args ? args.fwgrp : undefined;
            resourceInputs["loggrp"] = args ? args.loggrp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["netgrp"] = args ? args.netgrp : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["secfabgrp"] = args ? args.secfabgrp : undefined;
            resourceInputs["sysgrp"] = args ? args.sysgrp : undefined;
            resourceInputs["utmgrp"] = args ? args.utmgrp : undefined;
            resourceInputs["vpngrp"] = args ? args.vpngrp : undefined;
            resourceInputs["wanoptgrp"] = args ? args.wanoptgrp : undefined;
            resourceInputs["wifi"] = args ? args.wifi : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AdminProfiles.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdminProfiles resources.
 */
export interface AdminProfilesState {
    /**
     * Enable/disable overriding the global administrator idle timeout.
     */
    admintimeoutOverride?: pulumi.Input<string>;
    /**
     * Administrator access to Users and Devices.
     */
    authgrp?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * FortiView.
     */
    ftviewgrp?: pulumi.Input<string>;
    /**
     * Administrator access to the Firewall configuration.
     */
    fwgrp?: pulumi.Input<string>;
    /**
     * Administrator access to Logging and Reporting including viewing log messages.
     */
    loggrp?: pulumi.Input<string>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Network Configuration.
     */
    netgrp?: pulumi.Input<string>;
    /**
     * Scope of admin access.
     */
    scope?: pulumi.Input<string>;
    /**
     * Security Fabric.
     */
    secfabgrp?: pulumi.Input<string>;
    /**
     * System Configuration.
     */
    sysgrp?: pulumi.Input<string>;
    /**
     * Administrator access to Security Profiles.
     */
    utmgrp?: pulumi.Input<string>;
    /**
     * Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
     */
    vpngrp?: pulumi.Input<string>;
    /**
     * Administrator access to WAN Opt & Cache.
     */
    wanoptgrp?: pulumi.Input<string>;
    /**
     * Administrator access to the WiFi controller and Switch controller.
     */
    wifi?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AdminProfiles resource.
 */
export interface AdminProfilesArgs {
    /**
     * Enable/disable overriding the global administrator idle timeout.
     */
    admintimeoutOverride?: pulumi.Input<string>;
    /**
     * Administrator access to Users and Devices.
     */
    authgrp?: pulumi.Input<string>;
    /**
     * Comment.
     */
    comments?: pulumi.Input<string>;
    /**
     * FortiView.
     */
    ftviewgrp?: pulumi.Input<string>;
    /**
     * Administrator access to the Firewall configuration.
     */
    fwgrp?: pulumi.Input<string>;
    /**
     * Administrator access to Logging and Reporting including viewing log messages.
     */
    loggrp?: pulumi.Input<string>;
    /**
     * Profile name.
     */
    name?: pulumi.Input<string>;
    /**
     * Network Configuration.
     */
    netgrp?: pulumi.Input<string>;
    /**
     * Scope of admin access.
     */
    scope?: pulumi.Input<string>;
    /**
     * Security Fabric.
     */
    secfabgrp?: pulumi.Input<string>;
    /**
     * System Configuration.
     */
    sysgrp?: pulumi.Input<string>;
    /**
     * Administrator access to Security Profiles.
     */
    utmgrp?: pulumi.Input<string>;
    /**
     * Administrator access to IPsec, SSL, PPTP, and L2TP VPN.
     */
    vpngrp?: pulumi.Input<string>;
    /**
     * Administrator access to WAN Opt & Cache.
     */
    wanoptgrp?: pulumi.Input<string>;
    /**
     * Administrator access to the WiFi controller and Switch controller.
     */
    wifi?: pulumi.Input<string>;
}
