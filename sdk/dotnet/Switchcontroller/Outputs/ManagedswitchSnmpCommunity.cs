// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Switchcontroller.Outputs
{

    [OutputType]
    public sealed class ManagedswitchSnmpCommunity
    {
        /// <summary>
        /// SNMP notifications (traps) to send. Valid values: `cpu-high`, `mem-low`, `log-full`, `intf-ip`, `ent-conf-change`.
        /// </summary>
        public readonly string? Events;
        /// <summary>
        /// Configure IPv4 SNMP managers (hosts). The structure of `hosts` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ManagedswitchSnmpCommunityHost> Hosts;
        /// <summary>
        /// SNMP community ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// SNMP community name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// SNMP v1 query port (default = 161).
        /// </summary>
        public readonly int? QueryV1Port;
        /// <summary>
        /// Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? QueryV1Status;
        /// <summary>
        /// SNMP v2c query port (default = 161).
        /// </summary>
        public readonly int? QueryV2cPort;
        /// <summary>
        /// Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? QueryV2cStatus;
        /// <summary>
        /// Enable/disable this SNMP community. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// SNMP v2c trap local port (default = 162).
        /// </summary>
        public readonly int? TrapV1Lport;
        /// <summary>
        /// SNMP v2c trap remote port (default = 162).
        /// </summary>
        public readonly int? TrapV1Rport;
        /// <summary>
        /// Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? TrapV1Status;
        /// <summary>
        /// SNMP v2c trap local port (default = 162).
        /// </summary>
        public readonly int? TrapV2cLport;
        /// <summary>
        /// SNMP v2c trap remote port (default = 162).
        /// </summary>
        public readonly int? TrapV2cRport;
        /// <summary>
        /// Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.
        /// </summary>
        public readonly string? TrapV2cStatus;

        [OutputConstructor]
        private ManagedswitchSnmpCommunity(
            string? events,

            ImmutableArray<Outputs.ManagedswitchSnmpCommunityHost> hosts,

            int? id,

            string? name,

            int? queryV1Port,

            string? queryV1Status,

            int? queryV2cPort,

            string? queryV2cStatus,

            string? status,

            int? trapV1Lport,

            int? trapV1Rport,

            string? trapV1Status,

            int? trapV2cLport,

            int? trapV2cRport,

            string? trapV2cStatus)
        {
            Events = events;
            Hosts = hosts;
            Id = id;
            Name = name;
            QueryV1Port = queryV1Port;
            QueryV1Status = queryV1Status;
            QueryV2cPort = queryV2cPort;
            QueryV2cStatus = queryV2cStatus;
            Status = status;
            TrapV1Lport = trapV1Lport;
            TrapV1Rport = trapV1Rport;
            TrapV1Status = trapV1Status;
            TrapV2cLport = trapV2cLport;
            TrapV2cRport = trapV2cRport;
            TrapV2cStatus = trapV2cStatus;
        }
    }
}
