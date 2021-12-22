/******************************************************************************
 * Copyright (c)  2021 PingCAP, Inc.                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");            *
 * you may not use this file except in compliance with the License.           *
 * You may obtain a copy of the License at                                    *
 *                                                                            *
 * http://www.apache.org/licenses/LICENSE-2.0                                 *
 *                                                                            *
 * Unless required by applicable law or agreed to in writing, software        *
 * distributed under the License is distributed on an "AS IS" BASIS,          *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   *
 * See the License for the specific language governing permissions and        *
 * limitations under the License.                                             *
 *                                                                            *
 ******************************************************************************/

package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pingcap-inc/tiem/micro-api/controller"
	"github.com/pingcap-inc/tiem/micro-api/controller/cluster/backuprestore"
	changefeed2 "github.com/pingcap-inc/tiem/micro-api/controller/cluster/changefeed"
	logApi "github.com/pingcap-inc/tiem/micro-api/controller/cluster/log"
	clusterApi "github.com/pingcap-inc/tiem/micro-api/controller/cluster/management"
	parameterApi "github.com/pingcap-inc/tiem/micro-api/controller/cluster/parameter"
	"github.com/pingcap-inc/tiem/micro-api/controller/parametergroup"

	"github.com/pingcap-inc/tiem/micro-api/controller/datatransfer/importexport"
	"github.com/pingcap-inc/tiem/micro-api/controller/platform/specs"
	resourceApi "github.com/pingcap-inc/tiem/micro-api/controller/resource/hostresource"
	warehouseApi "github.com/pingcap-inc/tiem/micro-api/controller/resource/warehouse"
	flowtaskApi "github.com/pingcap-inc/tiem/micro-api/controller/task/flowtask"
	accountApi "github.com/pingcap-inc/tiem/micro-api/controller/user/account"
	idApi "github.com/pingcap-inc/tiem/micro-api/controller/user/identification"

	"github.com/pingcap-inc/tiem/micro-api/interceptor"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Route(g *gin.Engine) {
	// system check
	check := g.Group("/system")
	{
		check.GET("/check", controller.Hello)
	}

	// support swagger
	swagger := g.Group("/swagger")
	{
		swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// web
	web := g.Group("/web")
	{
		web.Use(interceptor.AccessLog(), gin.Recovery())
		// 替换成静态文件
		web.GET("/*any", controller.HelloPage)
	}

	// api
	apiV1 := g.Group("/api/v1")
	{
		apiV1.Use(interceptor.GinOpenTracing())
		apiV1.Use(interceptor.GinTraceIDHandler())
		apiV1.Use(interceptor.AccessLog(), gin.Recovery())

		user := apiV1.Group("/user")
		{
			user.POST("/login", idApi.Login)
			user.POST("/logout", idApi.Logout)
		}

		profile := user.Group("")
		{
			profile.Use(interceptor.VerifyIdentity)
			profile.Use(interceptor.AuditLog())
			profile.GET("/profile", accountApi.Profile)
		}

		cluster := apiV1.Group("/clusters")
		{
			cluster.Use(interceptor.VerifyIdentity)
			cluster.Use(interceptor.AuditLog())
			cluster.GET("/:clusterId", clusterApi.Detail)
			cluster.POST("/", clusterApi.Create)
			cluster.POST("/takeover", clusterApi.Takeover)
			cluster.POST("/preview", clusterApi.Preview)

			cluster.GET("/", clusterApi.Query)
			cluster.DELETE("/:clusterId", clusterApi.Delete)
			cluster.POST("/:clusterId/restart", clusterApi.Restart)
			cluster.POST("/:clusterId/stop", clusterApi.Stop)
			cluster.POST("/restore", backuprestore.Restore)
			cluster.GET("/:clusterId/dashboard", clusterApi.GetDashboardInfo)
			cluster.GET("/:clusterId/monitor", clusterApi.GetMonitorInfo)

			cluster.GET("/:clusterId/log", logApi.QueryClusterLog)

			// Scale cluster
			cluster.POST("/:clusterId/scale-out", clusterApi.ScaleOut)
			cluster.POST("/:clusterId/scale-in", clusterApi.ScaleIn)

			// Clone cluster
			cluster.POST("/clone", clusterApi.Clone)

			// Params
			cluster.GET("/:clusterId/params", parameterApi.QueryParameters)
			cluster.PUT("/:clusterId/params", parameterApi.UpdateParameters)
			//cluster.POST("/:clusterId/params/inspect", parameterApi.InspectParameters)

			// Backup Strategy
			cluster.GET("/:clusterId/strategy", backuprestore.GetBackupStrategy)
			cluster.PUT("/:clusterId/strategy", backuprestore.SaveBackupStrategy)
			// cluster.DELETE("/:clusterId/strategy", instanceapi.DeleteBackupStrategy)

			//Import and Export
			cluster.POST("/import", importexport.ImportData)
			cluster.POST("/export", importexport.ExportData)
			cluster.GET("/transport", importexport.QueryDataTransport)
			cluster.DELETE("/transport/:recordId", importexport.DeleteDataTransportRecord)
		}

		knowledge := apiV1.Group("/knowledges")
		{
			knowledge.GET("/", specs.ClusterKnowledge)
		}

		backup := apiV1.Group("/backups")
		{
			backup.Use(interceptor.VerifyIdentity)
			backup.Use(interceptor.AuditLog())

			backup.POST("/", backuprestore.Backup)
			backup.GET("/", backuprestore.QueryBackupRecords)
			backup.DELETE("/:backupId", backuprestore.DeleteBackup)
		}

		changeFeeds := apiV1.Group("/changefeeds")
		{
			changeFeeds.Use(interceptor.VerifyIdentity)
			changeFeeds.Use(interceptor.AuditLog())

			changeFeeds.POST("/", changefeed2.Create)
			changeFeeds.POST("/:changeFeedTaskId/pause", changefeed2.Pause)
			changeFeeds.POST("/:changeFeedTaskId/resume", changefeed2.Resume)
			changeFeeds.POST("/:changeFeedTaskId/update", changefeed2.Update)

			changeFeeds.DELETE("/:changeFeedTaskId", changefeed2.Delete)

			changeFeeds.GET("/:changeFeedTaskId", changefeed2.Detail)
			changeFeeds.GET("/", changefeed2.Query)
		}

		flowworks := apiV1.Group("/workflow")
		{
			flowworks.Use(interceptor.VerifyIdentity)
			flowworks.Use(interceptor.AuditLog())
			flowworks.GET("/", flowtaskApi.Query)
			flowworks.GET("/:workFlowId", flowtaskApi.Detail)
		}

		host := apiV1.Group("/resources")
		{
			host.Use(interceptor.VerifyIdentity)
			host.Use(interceptor.AuditLog())
			host.POST("hosts", resourceApi.ImportHosts)
			host.GET("hosts", resourceApi.QueryHosts)
			host.DELETE("hosts", resourceApi.RemoveHosts)
			host.GET("hosts-template", resourceApi.DownloadHostTemplateFile)
			host.GET("hierarchy", warehouseApi.GetHierarchy)
			host.GET("stocks", warehouseApi.GetStocks)
			host.PUT("host-reserved", resourceApi.UpdateHostReserved)
			host.PUT("host-status", resourceApi.UpdateHostStatus)
		}

		paramGroups := apiV1.Group("/param-groups")
		{
			paramGroups.Use(interceptor.VerifyIdentity)
			paramGroups.Use(interceptor.AuditLog())
			paramGroups.GET("/", parametergroup.Query)
			paramGroups.GET("/:paramGroupId", parametergroup.Detail)
			paramGroups.POST("/", parametergroup.Create)
			paramGroups.PUT("/:paramGroupId", parametergroup.Update)
			paramGroups.DELETE("/:paramGroupId", parametergroup.Delete)
			paramGroups.POST("/:paramGroupId/copy", parametergroup.Copy)
			paramGroups.POST("/:paramGroupId/apply", parametergroup.Apply)
		}
	}

}
