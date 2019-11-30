module Configuration
  module ProductResearch
    class Category
      attr_reader :name, :type, :urls

      def initialize(options)
        @name = options.fetch(:name)
        @type = options.fetch(:type)
        @urls = options.fetch(:urls).map { |u| URI.parse(u) }
      end
    end
  end
end
