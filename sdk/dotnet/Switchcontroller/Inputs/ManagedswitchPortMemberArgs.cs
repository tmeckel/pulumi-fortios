// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Inputs
{

    public sealed class ManagedswitchPortMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Interface name from available options.
        /// </summary>
        [Input("memberName")]
        public Input<string>? MemberName { get; set; }

        public ManagedswitchPortMemberArgs()
        {
        }
        public static new ManagedswitchPortMemberArgs Empty => new ManagedswitchPortMemberArgs();
    }
}
