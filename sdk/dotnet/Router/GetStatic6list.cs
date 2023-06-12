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
    public static class GetStatic6list
    {
        /// <summary>
        /// Provides a list of `fortios.router.Static6`.
        /// </summary>
        public static Task<GetStatic6listResult> InvokeAsync(GetStatic6listArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStatic6listResult>("fortios:router/getStatic6list:getStatic6list", args ?? new GetStatic6listArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.router.Static6`.
        /// </summary>
        public static Output<GetStatic6listResult> Invoke(GetStatic6listInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStatic6listResult>("fortios:router/getStatic6list:getStatic6list", args ?? new GetStatic6listInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStatic6listArgs : global::Pulumi.InvokeArgs
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

        public GetStatic6listArgs()
        {
        }
        public static new GetStatic6listArgs Empty => new GetStatic6listArgs();
    }

    public sealed class GetStatic6listInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetStatic6listInvokeArgs()
        {
        }
        public static new GetStatic6listInvokeArgs Empty => new GetStatic6listInvokeArgs();
    }


    [OutputType]
    public sealed class GetStatic6listResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.router.Static6`.
        /// </summary>
        public readonly ImmutableArray<int> SeqNumlists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetStatic6listResult(
            string? filter,

            string id,

            ImmutableArray<int> seqNumlists,

            string? vdomparam)
        {
            Filter = filter;
            Id = id;
            SeqNumlists = seqNumlists;
            Vdomparam = vdomparam;
        }
    }
}
