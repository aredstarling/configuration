module Configuration
  module ProductResearch
    class << self
      def find_marketplace_by_code(code)
        MARKETPLACES.find { |m| m.code == code }
      end
    end

    CONFIGURATION = begin
      path = File.join(__dir__, "/../../../product_research/#{ENV.fetch('APP_ENV')}.yml")
      Hanami::Utils::Hash.deep_symbolize(YAML.load_file(path))
    end.freeze

    MARKETPLACES = CONFIGURATION.fetch(:marketplaces).map { |c| Marketplace.new(c) }.freeze
  end
end
