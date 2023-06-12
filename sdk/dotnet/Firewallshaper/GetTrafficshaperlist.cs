// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewallshaper
{
    public static class GetTrafficshaperlist
    {
        /// <summary>
        /// Provides a list of `fortios.firewallshaper.Trafficshaper`.
        /// </summary>
        public static Task<GetTrafficshaperlistResult> InvokeAsync(GetTrafficshaperlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTrafficshaperlistResult>("fortios:firewallshaper/getTrafficshaperlist:getTrafficshaperlist", args ?? new GetTrafficshaperlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.firewallshaper.Trafficshaper`.
        /// </summary>
        public static Output<GetTrafficshaperlistResult> Invoke(GetTrafficshaperlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTrafficshaperlistResult>("fortios:firewallshaper/getTrafficshaperlist:getTrafficshaperlist", args ?? new GetTrafficshaperlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTrafficshaperlistArgs : global::Pulumi.InvokeArgs
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

        public GetTrafficshaperlistArgs()
        {
        }
        public static new GetTrafficshaperlistArgs Empty => new GetTrafficshaperlistArgs();
    }

    public sealed class GetTrafficshaperlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetTrafficshaperlistInvokeArgs()
        {
        }
        public static new GetTrafficshaperlistInvokeArgs Empty => new GetTrafficshaperlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetTrafficshaperlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.firewallshaper.Trafficshaper`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetTrafficshaperlistResult(
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
