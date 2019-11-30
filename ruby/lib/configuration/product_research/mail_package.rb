module Configuration
  module ProductResearch
    class MailPackage
      attr_reader :id, :name, :max_price, :max_dimensions, :max_weight, :max_length_girth

      def initialize(options)
        @id = options.fetch(:id)
        @name = options.fetch(:name)
        @max_price = options[:max_price]
        @max_dimensions = options[:max_dimensions]
        @max_weight = options[:max_weight]
        @max_length_girth = options[:max_length_girth]
      end
    end
  end
end
