// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall
{
    public static class GetDoSpolicylist
    {
        /// <summary>
        /// Provides a list of `fortios_firewall_DoSpolicy`.
        /// </summary>
        public static Task<GetDoSpolicylistResult> InvokeAsync(GetDoSpolicylistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDoSpolicylistResult>("fortios:firewall/getDoSpolicylist:getDoSpolicylist", args ?? new GetDoSpolicylistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios_firewall_DoSpolicy`.
        /// </summary>
        public static Output<GetDoSpolicylistResult> Invoke(GetDoSpolicylistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDoSpolicylistResult>("fortios:firewall/getDoSpolicylist:getDoSpolicylist", args ?? new GetDoSpolicylistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDoSpolicylistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public string? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetDoSpolicylistArgs()
        {
        }
        public static new GetDoSpolicylistArgs Empty => new GetDoSpolicylistArgs();
    }

    public sealed class GetDoSpolicylistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter used to scope the list. See Filter results of datasource.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetDoSpolicylistInvokeArgs()
        {
        }
        public static new GetDoSpolicylistInvokeArgs Empty => new GetDoSpolicylistInvokeArgs();
    }


    [OutputType]
    public sealed class GetDoSpolicylistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios_firewall_DoSpolicy`.
        /// </summary>
        public readonly ImmutableArray<int> Policyidlists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetDoSpolicylistResult(
            string? filter,

            string id,

            ImmutableArray<int> policyidlists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            Policyidlists = policyidlists;
            Vdomparam = vdomparam;
        }
    }
}
