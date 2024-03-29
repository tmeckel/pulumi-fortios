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
    public static class GetAccesslist6list
    {
        /// <summary>
        /// Provides a list of `fortios.router.Accesslist6`.
        /// </summary>
        public static Task<GetAccesslist6listResult> InvokeAsync(GetAccesslist6listArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccesslist6listResult>("fortios:router/getAccesslist6list:getAccesslist6list", args ?? new GetAccesslist6listArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a list of `fortios.router.Accesslist6`.
        /// </summary>
        public static Output<GetAccesslist6listResult> Invoke(GetAccesslist6listInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccesslist6listResult>("fortios:router/getAccesslist6list:getAccesslist6list", args ?? new GetAccesslist6listInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccesslist6listArgs : global::Pulumi.InvokeArgs
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

        public GetAccesslist6listArgs()
        {
        }
        public static new GetAccesslist6listArgs Empty => new GetAccesslist6listArgs();
    }

    public sealed class GetAccesslist6listInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetAccesslist6listInvokeArgs()
        {
        }
        public static new GetAccesslist6listInvokeArgs Empty => new GetAccesslist6listInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccesslist6listResult
    {
        public readonly string? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of the `fortios.router.Accesslist6`.
        /// </summary>
        public readonly ImmutableArray<string> Namelists;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetAccesslist6listResult(
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
