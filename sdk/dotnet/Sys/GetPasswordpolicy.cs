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
    public static class GetPasswordpolicy
    {
        /// <summary>
        /// Use this data source to get information on fortios system passwordpolicy
        /// </summary>
        public static Task<GetPasswordpolicyResult> InvokeAsync(GetPasswordpolicyArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPasswordpolicyResult>("fortios:sys/getPasswordpolicy:getPasswordpolicy", args ?? new GetPasswordpolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system passwordpolicy
        /// </summary>
        public static Output<GetPasswordpolicyResult> Invoke(GetPasswordpolicyInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPasswordpolicyResult>("fortios:sys/getPasswordpolicy:getPasswordpolicy", args ?? new GetPasswordpolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPasswordpolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetPasswordpolicyArgs()
        {
        }
        public static new GetPasswordpolicyArgs Empty => new GetPasswordpolicyArgs();
    }

    public sealed class GetPasswordpolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetPasswordpolicyInvokeArgs()
        {
        }
        public static new GetPasswordpolicyInvokeArgs Empty => new GetPasswordpolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetPasswordpolicyResult
    {
        /// <summary>
        /// Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space.
        /// </summary>
        public readonly string ApplyTo;
        /// <summary>
        /// Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled).
        /// </summary>
        public readonly string Change4Characters;
        /// <summary>
        /// Number of days after which passwords expire (1 - 999 days, default = 90).
        /// </summary>
        public readonly int ExpireDay;
        /// <summary>
        /// Enable/disable password expiration.
        /// </summary>
        public readonly string ExpireStatus;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Minimum number of unique characters in new password which do not exist in old password (This attribute overrides reuse-password if both are enabled).
        /// </summary>
        public readonly int MinChangeCharacters;
        /// <summary>
        /// Minimum number of lowercase characters in password (0 - 128, default = 0).
        /// </summary>
        public readonly int MinLowerCaseLetter;
        /// <summary>
        /// Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).
        /// </summary>
        public readonly int MinNonAlphanumeric;
        /// <summary>
        /// Minimum number of numeric characters in password (0 - 128, default = 0).
        /// </summary>
        public readonly int MinNumber;
        /// <summary>
        /// Minimum number of uppercase characters in password (0 - 128, default = 0).
        /// </summary>
        public readonly int MinUpperCaseLetter;
        /// <summary>
        /// Minimum password length (8 - 128, default = 8).
        /// </summary>
        public readonly int MinimumLength;
        /// <summary>
        /// Enable/disable reusing of password (if both reuse-password and change-4-characters are enabled, change-4-characters overrides).
        /// </summary>
        public readonly string ReusePassword;
        /// <summary>
        /// Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.
        /// </summary>
        public readonly string Status;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetPasswordpolicyResult(
            string applyTo,

            string change4Characters,

            int expireDay,

            string expireStatus,

            string id,

            int minChangeCharacters,

            int minLowerCaseLetter,

            int minNonAlphanumeric,

            int minNumber,

            int minUpperCaseLetter,

            int minimumLength,

            string reusePassword,

            string status,

            string? vdomparam)
        {
            ApplyTo = applyTo;
            Change4Characters = change4Characters;
            ExpireDay = expireDay;
            ExpireStatus = expireStatus;
            Id = id;
            MinChangeCharacters = minChangeCharacters;
            MinLowerCaseLetter = minLowerCaseLetter;
            MinNonAlphanumeric = minNonAlphanumeric;
            MinNumber = minNumber;
            MinUpperCaseLetter = minUpperCaseLetter;
            MinimumLength = minimumLength;
            ReusePassword = reusePassword;
            Status = status;
            Vdomparam = vdomparam;
        }
    }
}
