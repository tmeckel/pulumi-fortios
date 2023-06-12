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
    public static class GetAddrgrp
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall addrgrp
        /// </summary>
        public static Task<GetAddrgrpResult> InvokeAsync(GetAddrgrpArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAddrgrpResult>("fortios:firewall/getAddrgrp:getAddrgrp", args ?? new GetAddrgrpArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall addrgrp
        /// </summary>
        public static Output<GetAddrgrpResult> Invoke(GetAddrgrpInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAddrgrpResult>("fortios:firewall/getAddrgrp:getAddrgrp", args ?? new GetAddrgrpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAddrgrpArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall addrgrp.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetAddrgrpArgs()
        {
        }
        public static new GetAddrgrpArgs Empty => new GetAddrgrpArgs();
    }

    public sealed class GetAddrgrpInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall addrgrp.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetAddrgrpInvokeArgs()
        {
        }
        public static new GetAddrgrpInvokeArgs Empty => new GetAddrgrpInvokeArgs();
    }


    [OutputType]
    public sealed class GetAddrgrpResult
    {
        /// <summary>
        /// Enable/disable use of this group in the static route configuration.
        /// </summary>
        public readonly string AllowRouting;
        /// <summary>
        /// Tag category.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        public readonly int Color;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// Enable/disable address exclusion.
        /// </summary>
        public readonly string Exclude;
        /// <summary>
        /// Address exclusion member. The structure of `exclude_member` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAddrgrpExcludeMemberResult> ExcludeMembers;
        /// <summary>
        /// Security Fabric global object setting.
        /// </summary>
        public readonly string FabricObject;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Address objects contained within the group. The structure of `member` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAddrgrpMemberResult> Members;
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAddrgrpTaggingResult> Taggings;
        /// <summary>
        /// Address group type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
        /// </summary>
        public readonly string Uuid;
        public readonly string? Vdomparam;
        /// <summary>
        /// Enable/disable address visibility in the GUI.
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetAddrgrpResult(
            string allowRouting,

            string category,

            int color,

            string comment,

            string exclude,

            ImmutableArray<Outputs.GetAddrgrpExcludeMemberResult> excludeMembers,

            string fabricObject,

            string id,

            ImmutableArray<Outputs.GetAddrgrpMemberResult> members,

            string name,

            ImmutableArray<Outputs.GetAddrgrpTaggingResult> taggings,

            string type,

            string uuid,

            string? vdomparam,

            string visibility)
        {
            AllowRouting = allowRouting;
            Category = category;
            Color = color;
            Comment = comment;
            Exclude = exclude;
            ExcludeMembers = excludeMembers;
            FabricObject = fabricObject;
            Id = id;
            Members = members;
            Name = name;
            Taggings = taggings;
            Type = type;
            Uuid = uuid;
            Vdomparam = vdomparam;
            Visibility = visibility;
        }
    }
}