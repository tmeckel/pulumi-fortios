// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Application.Inputs
{

    public sealed class ListDefaultNetworkServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entry ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Port number.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Network protocols. Valid values: `http`, `ssh`, `telnet`, `ftp`, `dns`, `smtp`, `pop3`, `imap`, `snmp`, `nntp`, `https`.
        /// </summary>
        [Input("services")]
        public Input<string>? Services { get; set; }

        /// <summary>
        /// Action for protocols not white listed under selected port. Valid values: `pass`, `monitor`, `block`.
        /// </summary>
        [Input("violationAction")]
        public Input<string>? ViolationAction { get; set; }

        public ListDefaultNetworkServiceArgs()
        {
        }
        public static new ListDefaultNetworkServiceArgs Empty => new ListDefaultNetworkServiceArgs();
    }
}