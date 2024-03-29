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

    public sealed class ManagedswitchSnmpUserGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication protocol.
        /// </summary>
        [Input("authProto")]
        public Input<string>? AuthProto { get; set; }

        [Input("authPwd")]
        private Input<string>? _authPwd;

        /// <summary>
        /// Password for authentication protocol.
        /// </summary>
        public Input<string>? AuthPwd
        {
            get => _authPwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _authPwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// SNMP user name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Privacy (encryption) protocol.
        /// </summary>
        [Input("privProto")]
        public Input<string>? PrivProto { get; set; }

        [Input("privPwd")]
        private Input<string>? _privPwd;

        /// <summary>
        /// Password for privacy (encryption) protocol.
        /// </summary>
        public Input<string>? PrivPwd
        {
            get => _privPwd;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privPwd = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.
        /// </summary>
        [Input("queries")]
        public Input<string>? Queries { get; set; }

        /// <summary>
        /// SNMPv3 query port (default = 161).
        /// </summary>
        [Input("queryPort")]
        public Input<int>? QueryPort { get; set; }

        /// <summary>
        /// Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.
        /// </summary>
        [Input("securityLevel")]
        public Input<string>? SecurityLevel { get; set; }

        public ManagedswitchSnmpUserGetArgs()
        {
        }
        public static new ManagedswitchSnmpUserGetArgs Empty => new ManagedswitchSnmpUserGetArgs();
    }
}
