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
    public static class GetInternetservicecustomgrouplist
    {
        /// <summary>
        /// Provides a list of `fortios.firewall.Internetservicecustomgroup`.
        /// </summary>
        public static Task<GetInternetservicecustomgrouplistResult> InvokeAsync(GetInternetservicecustomgrouplistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInternetservicecustomgrouplistResult>("fortios:firewall/getInternetservicecustomgrouplist:getInternetservicecustomgrouplist", args ?? new GetInternetservicecustomgrouplistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.firewall.Internetservicecustomgroup`.
        /// </summary>
        public static Output<GetInternetservicecustomgrouplistResult> Invoke(GetInternetservicecustomgrouplistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInternetservicecustomgrouplistResult>("fortios:firewall/getInternetservicecustomgrouplist:getInternetservicecustomgrouplist", args ?? new GetInternetservicecustomgrouplistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInternetservicecustomgrouplistArgs : global::Pulumi.InvokeArgs
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

        public GetInternetservicecustomgrouplistArgs()
        {
        }
        public static new GetInternetservicecustomgrouplistArgs Empty => new GetInternetservicecustomgrouplistArgs();
    }

    public sealed class GetInternetservicecustomgrouplistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetInternetservicecustomgrouplistInvokeArgs()
        {
        }
        public static new GetInternetservicecustomgrouplistInvokeArgs Empty => new GetInternetservicecustomgrouplistInvokeArgs();
    }


    [OutputType]
    public sealed class GetInternetservicecustomgrouplistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.firewall.Internetservicecustomgroup`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetInternetservicecustomgrouplistResult(
            string? filter,

            string id,

            ImmutableArray<string> namelists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            Namelists = namelists;
            Vdomparam = vdomparam;
        }
    }
}
