// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Sys
{
    /// <summary>
    /// Configure the password policy for guest administrators.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Fortios = Pulumiverse.Fortios;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var trname = new Fortios.Sys.Passwordpolicyguestadmin("trname", new()
    ///     {
    ///         ApplyTo = "guest-admin-password",
    ///         Change4Characters = "disable",
    ///         ExpireDay = 90,
    ///         ExpireStatus = "disable",
    ///         MinLowerCaseLetter = 0,
    ///         MinNonAlphanumeric = 0,
    ///         MinNumber = 0,
    ///         MinUpperCaseLetter = 0,
    ///         MinimumLength = 8,
    ///         ReusePassword = "enable",
    ///         Status = "disable",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// System PasswordPolicyGuestAdmin can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/passwordpolicyguestadmin:Passwordpolicyguestadmin labelname SystemPasswordPolicyGuestAdmin
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:sys/passwordpolicyguestadmin:Passwordpolicyguestadmin labelname SystemPasswordPolicyGuestAdmin
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:sys/passwordpolicyguestadmin:Passwordpolicyguestadmin")]
    public partial class Passwordpolicyguestadmin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
        /// </summary>
        [Output("applyTo")]
        public Output<string> ApplyTo { get; private set; } = null!;

        /// <summary>
        /// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("change4Characters")]
        public Output<string> Change4Characters { get; private set; } = null!;

        /// <summary>
        /// Number of days after which passwords expire (1 - 999 days, default = 90).
        /// </summary>
        [Output("expireDay")]
        public Output<int> ExpireDay { get; private set; } = null!;

        /// <summary>
        /// Enable/disable password expiration. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("expireStatus")]
        public Output<string> ExpireStatus { get; private set; } = null!;

        /// <summary>
        /// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
        /// </summary>
        [Output("minChangeCharacters")]
        public Output<int> MinChangeCharacters { get; private set; } = null!;

        /// <summary>
        /// Minimum number of lowercase characters in password (0 - 128, default = 0).
        /// </summary>
        [Output("minLowerCaseLetter")]
        public Output<int> MinLowerCaseLetter { get; private set; } = null!;

        /// <summary>
        /// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        /// </summary>
        [Output("minNonAlphanumeric")]
        public Output<int> MinNonAlphanumeric { get; private set; } = null!;

        /// <summary>
        /// Minimum number of numeric characters in password (0 - 128, default = 0).
        /// </summary>
        [Output("minNumber")]
        public Output<int> MinNumber { get; private set; } = null!;

        /// <summary>
        /// Minimum number of uppercase characters in password (0 - 128, default = 0).
        /// </summary>
        [Output("minUpperCaseLetter")]
        public Output<int> MinUpperCaseLetter { get; private set; } = null!;

        /// <summary>
        /// Minimum password length (8 - 128, default = 8).
        /// </summary>
        [Output("minimumLength")]
        public Output<int> MinimumLength { get; private set; } = null!;

        /// <summary>
        /// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
        /// </summary>
        [Output("reusePassword")]
        public Output<string> ReusePassword { get; private set; } = null!;

        /// <summary>
        /// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a Passwordpolicyguestadmin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Passwordpolicyguestadmin(string name, PasswordpolicyguestadminArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:sys/passwordpolicyguestadmin:Passwordpolicyguestadmin", name, args ?? new PasswordpolicyguestadminArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Passwordpolicyguestadmin(string name, Input<string> id, PasswordpolicyguestadminState? state = null, CustomResourceOptions? options = null)
            : base("fortios:sys/passwordpolicyguestadmin:Passwordpolicyguestadmin", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Passwordpolicyguestadmin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Passwordpolicyguestadmin Get(string name, Input<string> id, PasswordpolicyguestadminState? state = null, CustomResourceOptions? options = null)
        {
            return new Passwordpolicyguestadmin(name, id, state, options);
        }
    }

    public sealed class PasswordpolicyguestadminArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
        /// </summary>
        [Input("applyTo")]
        public Input<string>? ApplyTo { get; set; }

        /// <summary>
        /// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("change4Characters")]
        public Input<string>? Change4Characters { get; set; }

        /// <summary>
        /// Number of days after which passwords expire (1 - 999 days, default = 90).
        /// </summary>
        [Input("expireDay")]
        public Input<int>? ExpireDay { get; set; }

        /// <summary>
        /// Enable/disable password expiration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("expireStatus")]
        public Input<string>? ExpireStatus { get; set; }

        /// <summary>
        /// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
        /// </summary>
        [Input("minChangeCharacters")]
        public Input<int>? MinChangeCharacters { get; set; }

        /// <summary>
        /// Minimum number of lowercase characters in password (0 - 128, default = 0).
        /// </summary>
        [Input("minLowerCaseLetter")]
        public Input<int>? MinLowerCaseLetter { get; set; }

        /// <summary>
        /// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        /// </summary>
        [Input("minNonAlphanumeric")]
        public Input<int>? MinNonAlphanumeric { get; set; }

        /// <summary>
        /// Minimum number of numeric characters in password (0 - 128, default = 0).
        /// </summary>
        [Input("minNumber")]
        public Input<int>? MinNumber { get; set; }

        /// <summary>
        /// Minimum number of uppercase characters in password (0 - 128, default = 0).
        /// </summary>
        [Input("minUpperCaseLetter")]
        public Input<int>? MinUpperCaseLetter { get; set; }

        /// <summary>
        /// Minimum password length (8 - 128, default = 8).
        /// </summary>
        [Input("minimumLength")]
        public Input<int>? MinimumLength { get; set; }

        /// <summary>
        /// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("reusePassword")]
        public Input<string>? ReusePassword { get; set; }

        /// <summary>
        /// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PasswordpolicyguestadminArgs()
        {
        }
        public static new PasswordpolicyguestadminArgs Empty => new PasswordpolicyguestadminArgs();
    }

    public sealed class PasswordpolicyguestadminState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Guest administrator to which this password policy applies. Valid values: `guest-admin-password`.
        /// </summary>
        [Input("applyTo")]
        public Input<string>? ApplyTo { get; set; }

        /// <summary>
        /// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("change4Characters")]
        public Input<string>? Change4Characters { get; set; }

        /// <summary>
        /// Number of days after which passwords expire (1 - 999 days, default = 90).
        /// </summary>
        [Input("expireDay")]
        public Input<int>? ExpireDay { get; set; }

        /// <summary>
        /// Enable/disable password expiration. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("expireStatus")]
        public Input<string>? ExpireStatus { get; set; }

        /// <summary>
        /// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
        /// </summary>
        [Input("minChangeCharacters")]
        public Input<int>? MinChangeCharacters { get; set; }

        /// <summary>
        /// Minimum number of lowercase characters in password (0 - 128, default = 0).
        /// </summary>
        [Input("minLowerCaseLetter")]
        public Input<int>? MinLowerCaseLetter { get; set; }

        /// <summary>
        /// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        /// </summary>
        [Input("minNonAlphanumeric")]
        public Input<int>? MinNonAlphanumeric { get; set; }

        /// <summary>
        /// Minimum number of numeric characters in password (0 - 128, default = 0).
        /// </summary>
        [Input("minNumber")]
        public Input<int>? MinNumber { get; set; }

        /// <summary>
        /// Minimum number of uppercase characters in password (0 - 128, default = 0).
        /// </summary>
        [Input("minUpperCaseLetter")]
        public Input<int>? MinUpperCaseLetter { get; set; }

        /// <summary>
        /// Minimum password length (8 - 128, default = 8).
        /// </summary>
        [Input("minimumLength")]
        public Input<int>? MinimumLength { get; set; }

        /// <summary>
        /// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides). Valid values: `enable`, `disable`.
        /// </summary>
        [Input("reusePassword")]
        public Input<string>? ReusePassword { get; set; }

        /// <summary>
        /// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public PasswordpolicyguestadminState()
        {
        }
        public static new PasswordpolicyguestadminState Empty => new PasswordpolicyguestadminState();
    }
}