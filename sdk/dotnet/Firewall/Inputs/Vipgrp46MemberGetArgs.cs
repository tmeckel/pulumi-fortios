// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Firewall.Inputs
{

    public sealed class Vipgrp46MemberGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VIP46 name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Vipgrp46MemberGetArgs()
        {
        }
        public static new Vipgrp46MemberGetArgs Empty => new Vipgrp46MemberGetArgs();
    }
}
