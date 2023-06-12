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
    public static class GetInternetserviceextensionlist
    {
        /// <summary>
        /// Provides a list of `fortios.firewall.Internetserviceextension`.
        /// </summary>
        public static Task<GetInternetserviceextensionlistResult> InvokeAsync(GetInternetserviceextensionlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInternetserviceextensionlistResult>("fortios:firewall/getInternetserviceextensionlist:getInternetserviceextensionlist", args ?? new GetInternetserviceextensionlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.firewall.Internetserviceextension`.
        /// </summary>
        public static Output<GetInternetserviceextensionlistResult> Invoke(GetInternetserviceextensionlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInternetserviceextensionlistResult>("fortios:firewall/getInternetserviceextensionlist:getInternetserviceextensionlist", args ?? new GetInternetserviceextensionlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInternetserviceextensionlistArgs : global::Pulumi.InvokeArgs
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

        public GetInternetserviceextensionlistArgs()
        {
        }
        public static new GetInternetserviceextensionlistArgs Empty => new GetInternetserviceextensionlistArgs();
    }

    public sealed class GetInternetserviceextensionlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetInternetserviceextensionlistInvokeArgs()
        {
        }
        public static new GetInternetserviceextensionlistInvokeArgs Empty => new GetInternetserviceextensionlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetInternetserviceextensionlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// A list of the `fortios.firewall.Internetserviceextension`.
        /// </summary>
        public readonly ImmutableArray<int> Fosidlists;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetInternetserviceextensionlistResult(
            string? filter,

            ImmutableArray<int> fosidlists,

            string id,

            string? vdomparam)
        {
            Filter = filter;
            Fosidlists = fosidlists;
            Id = id;
            Vdomparam = vdomparam;
        }
    }
}