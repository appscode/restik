/*
Copyright The Stash Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package misc

import (
	"stash.appscode.dev/stash/apis"
	"stash.appscode.dev/stash/apis/stash/v1beta1"
	"stash.appscode.dev/stash/pkg/eventer"
	"stash.appscode.dev/stash/test/e2e/framework"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
)

var _ = Describe("Failure Handling", func() {

	var f *framework.Invocation

	BeforeEach(func() {
		f = framework.NewInvocation()
	})

	JustAfterEach(func() {
		f.PrintDebugInfoOnFailure()
	})

	AfterEach(func() {
		err := f.CleanupTestResources()
		Expect(err).NotTo(HaveOccurred())
	})

	Context("CronJob Creation Failure", func() {
		It("should write CronJob creation failure event to the BackupConfiguration", func() {
			// Deploy a Deployment
			deployment, err := f.DeployDeployment(framework.SourceDeployment, int32(1), framework.SourceVolume)
			Expect(err).NotTo(HaveOccurred())

			// Generate Sample Data
			_, err = f.GenerateSampleData(deployment.ObjectMeta, apis.KindDeployment)
			Expect(err).NotTo(HaveOccurred())

			// Setup a Minio Repository
			repo, err := f.SetupMinioRepository()
			Expect(err).NotTo(HaveOccurred())
			f.AppendToCleanupList(repo)

			// Setup workload Backup
			backupConfig, err := f.CreateBackupConfigForWorkload(deployment.ObjectMeta, repo, apis.KindDeployment, func(bc *v1beta1.BackupConfiguration) {
				bc.Spec.Schedule = "*/3 * * * * *"
			})
			Expect(err).NotTo(HaveOccurred())

			By("Verifying that CronJob creation failure event has been written")
			f.EventuallyEventWritten(
				backupConfig.ObjectMeta,
				v1beta1.ResourceKindBackupConfiguration,
				v1.EventTypeWarning,
				eventer.EventReasonCronJobCreationFailed,
			).Should(BeTrue())
		})
	})
})