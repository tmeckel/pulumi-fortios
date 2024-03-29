// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Router
{
    public static class GetAuthpathlist
    {
        /// <summary>
        /// Provides a list of `fortios.router.Authpath`.
        /// </summary>
        public static Task<GetAuthpathlistResult> InvokeAsync(GetAuthpathlistArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthpathlistResult>("fortios:router/getAuthpathlist:getAuthpathlist", args ?? new GetAuthpathlistArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.router.Authpath`.
        /// </summary>
        public static Output<GetAuthpathlistResult> Invoke(GetAuthpathlistInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthpathlistResult>("fortios:router/getAuthpathlist:getAuthpathlist", args ?? new GetAuthpathlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthpathlistArgs : global::Pulumi.InvokeArgs
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

        public GetAuthpathlistArgs()
        {
        }
        public static new GetAuthpathlistArgs Empty => new GetAuthpathlistArgs();
    }

    public sealed class GetAuthpathlistInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetAuthpathlistInvokeArgs()
        {
        }
        public static new GetAuthpathlistInvokeArgs Empty => new GetAuthpathlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthpathlistResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.router.Authpath`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetAuthpathlistResult(
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
