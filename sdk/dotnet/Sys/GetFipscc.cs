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
    public static class GetFipscc
    {
        /// <summary>
        /// Use this data source to get information on fortios system fipscc
        /// </summary>
        public static Task<GetFipsccResult> InvokeAsync(GetFipsccArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFipsccResult>("fortios:sys/getFipscc:getFipscc", args ?? new GetFipsccArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios system fipscc
        /// </summary>
        public static Output<GetFipsccResult> Invoke(GetFipsccInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFipsccResult>("fortios:sys/getFipscc:getFipscc", args ?? new GetFipsccInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFipsccArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFipsccArgs()
        {
        }
        public static new GetFipsccArgs Empty => new GetFipsccArgs();
    }

    public sealed class GetFipsccInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFipsccInvokeArgs()
        {
        }
        public static new GetFipsccInvokeArgs Empty => new GetFipsccInvokeArgs();
    }


    [OutputType]
    public sealed class GetFipsccResult
    {
        /// <summary>
        /// Enable/disable/dynamic entropy token.
        /// </summary>
        public readonly string EntropyToken;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable/disable self tests after key generation.
        /// </summary>
        public readonly string KeyGenerationSelfTest;
        /// <summary>
        /// Self test period.
        /// </summary>
        public readonly int SelfTestPeriod;
        /// <summary>
        /// Enable/disable FIPS-CC mode.
        /// </summary>
        public readonly string Status;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFipsccResult(
            string entropyToken,

            string id,

            string keyGenerationSelfTest,

            int selfTestPeriod,

            string status,

            string? vdomparam)
        {
            EntropyToken = entropyToken;
            Id = id;
            KeyGenerationSelfTest = keyGenerationSelfTest;
            SelfTestPeriod = selfTestPeriod;
            Status = status;
            Vdomparam = vdomparam;
        }
    }
}
