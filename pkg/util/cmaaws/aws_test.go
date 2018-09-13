package cmaaws_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/cmaaws"
)

var _ = Describe("CMAAWS API Functions", func() {
	var (
		goodClient cmaaws.ClientInterface
		badClient  cmaaws.ClientInterface
	)
	BeforeEach(func() {
		goodClient = &cmaaws.Client{}
		goodClient.SetClient(&MockGoodCMAAWS{})
		badClient = &cmaaws.Client{}
		badClient.SetClient(&MockBadCMAAWS{})
	})
	Context("when creating a cluster", func() {
		var (
			err      error
			response cmaaws.CreateClusterOutput
		)
		Context("when the apiserver says it is successful", func() {
			BeforeEach(func() {
				response, err = goodClient.CreateCluster(cmaaws.CreateClusterInput{})
			})
			It("return the apiserver message", func() {
				Expect(response.Cluster.ID).To(Equal(MockGoodCreateClusterReply.Cluster.Id))
				Expect(response.Cluster.Name).To(Equal(MockGoodCreateClusterReply.Cluster.Name))
				Expect(response.Cluster.Status).To(Equal(MockGoodCreateClusterReply.Cluster.Status))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.CreateCluster(cmaaws.CreateClusterInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(cmaaws.CreateClusterOutput{}))
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
			It("return apiserver error message", func() {
				Expect(err).To(Equal(MockBadCreateClusterError))
			})
		})
	})

	Context("when getting a cluster", func() {
		var (
			err      error
			response cmaaws.GetClusterOutput
		)
		Context("when the apiserver says it is successful", func() {
			BeforeEach(func() {
				response, err = goodClient.GetCluster(cmaaws.GetClusterInput{})
			})
			It("return the apiserver message", func() {
				Expect(response.Cluster.ID).To(Equal(MockGoodGetClusterReply.Cluster.Id))
				Expect(response.Cluster.Name).To(Equal(MockGoodGetClusterReply.Cluster.Name))
				Expect(response.Cluster.Status).To(Equal(MockGoodGetClusterReply.Cluster.Status))
				Expect(response.Cluster.Kubeconfig).To(Equal(MockGoodGetClusterReply.Cluster.Kubeconfig))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.GetCluster(cmaaws.GetClusterInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(cmaaws.GetClusterOutput{}))
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
			It("return apiserver error message", func() {
				Expect(err).To(Equal(MockBadGetClusterError))
			})
		})
	})

	Context("when deleting a cluster", func() {
		var (
			err      error
			response cmaaws.DeleteClusterOutput
		)
		Context("when the apiserver says it is successful", func() {
			BeforeEach(func() {
				response, err = goodClient.DeleteCluster(cmaaws.DeleteClusterInput{})
			})
			It("return the apiserver message", func() {
				Expect(response.Status).To(Equal(MockGoodDeleteClusterReply.Status))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.DeleteCluster(cmaaws.DeleteClusterInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(cmaaws.DeleteClusterOutput{}))
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
			It("return apiserver error message", func() {
				Expect(err).To(Equal(MockBadDeleteClusterError))
			})
		})
	})

	Context("when getting a cluster list", func() {
		var (
			err      error
			response cmaaws.ListClusterOutput
		)
		Context("when the apiserver says it is successful", func() {
			BeforeEach(func() {
				response, err = goodClient.ListClusters(cmaaws.ListClusterInput{})
			})
			It("return the apiserver message", func() {
				Expect(response.Clusters).To(HaveLen(2))
				Expect(response.Clusters[0].Status).To(BeEquivalentTo(MockGoodGetClusterListReply.Clusters[0].Status))
				Expect(response.Clusters[1].ID).To(BeEquivalentTo(MockGoodGetClusterListReply.Clusters[1].Id))
				Expect(response.Clusters[0].Name).To(BeEquivalentTo(MockGoodGetClusterListReply.Clusters[0].Name))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.ListClusters(cmaaws.ListClusterInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(cmaaws.ListClusterOutput{}))
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
			It("return apiserver error message", func() {
				Expect(err).To(Equal(MockBadGetClusterListError))
			})
		})
	})
})
